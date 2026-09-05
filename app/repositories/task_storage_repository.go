package repositories

import (
	"context"
	"errors"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories/storage"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	TaskAssigneeNotFoundError = errors.New("tasks assignee not found")
)

type TaskStorageRepositoryFilterQuery struct {
	BoardId   *int `json:"board_id" form:"board_id"`
	ProjectId *int `json:"project_id" form:"project_id"`
}

type TaskStorageRepository interface {
	Filter(ctx context.Context, boardId int) ([]models.Task, error)
	FindByID(ctx context.Context, id int) (*models.Task, error)

	// Save creates or updates the entire cluster transactionally
	Save(ctx context.Context, task *models.Task) error

	// Explicit domain action optimized to bypass heavy loads when necessary
	AssignUser(ctx context.Context, taskID, userID int) error
	UnassignUser(ctx context.Context, taskID, userID int) error
	MovePosition(ctx context.Context, taskID, boardID, columnID, pos int) error
	Complete(ctx context.Context, taskID int) error
}

type gormTaskStorageRepository struct {
	db *gorm.DB
}

var _ TaskStorageRepository = (*gormTaskStorageRepository)(nil)

func NewGormTaskStorageRepository(db *gorm.DB) TaskStorageRepository {
	return &gormTaskStorageRepository{db: db}
}

// AssignUser implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) AssignUser(ctx context.Context, taskID, userID int) error {
	return gorm.G[storage.TaskAssigneeRecord](s.db, clause.OnConflict{DoNothing: true}).
		Create(ctx, &storage.TaskAssigneeRecord{
			TaskID: taskID,
			UserID: userID,
		})
}

// AssignUser implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) UnassignUser(ctx context.Context, taskID, userID int) error {
	rowsAffected, err := gorm.G[storage.TaskAssigneeRecord](s.db, clause.OnConflict{DoNothing: true}).
		Where("task_id = ?", taskID).
		Where("user_id = ?", userID).
		Delete(ctx)

	if rowsAffected == 0 {
		return TaskAssigneeNotFoundError
	}

	return err
}

// Complete implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Complete(ctx context.Context, taskID int) error {
	rows, err := gorm.G[storage.TaskRecord](s.db).
		Where("id = ?", taskID).
		Select("completed", "updated_at").
		Updates(ctx, storage.TaskRecord{Completed: true, UpdatedAt: time.Now()})
	if err != nil {
		return err
	}
	if rows == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// FindByID implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) FindByID(ctx context.Context, id int) (*models.Task, error) {
	task, err := gorm.G[storage.TaskRecord](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := gorm.G[storage.TaskOwnerRecord](s.db).Where("task_id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	assignee, err := gorm.G[storage.TaskAssigneeRecord](s.db).Where("task_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	boards, err := gorm.G[storage.BoardTaskRecord](s.db).Where("task_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	t := s.toDomainModels([]storage.TaskRecord{task}, []storage.TaskOwnerRecord{owner}, assignee, boards)
	if len(t) == 0 {
		return &models.Task{}, nil
	}
	return &t[0], nil
}

// MovePosition implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) MovePosition(ctx context.Context, taskID, boardID, columnID, pos int) error {
	rows, err := gorm.G[map[string]any](s.db).
		Table("board_tasks").
		Where("task_id = ? AND board_id = ?", taskID, boardID).
		Select("column_id", "position").
		Updates(ctx, map[string]any{"board_id": boardID, "task_id": taskID, "column_id": columnID, "position": pos})
	if err != nil {
		return err
	}
	if rows == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Save implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Save(ctx context.Context, task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.saveTaskRecord(ctx, tx, task); err != nil {
			return err
		}
		if err := s.saveOwner(ctx, tx, task); err != nil {
			return err
		}
		if err := s.saveAssignees(ctx, tx, task); err != nil {
			return err
		}
		return s.saveBoards(ctx, tx, task)
	})
}

// Filter implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Filter(ctx context.Context, boardId int) ([]models.Task, error) {
	tasks, err := gorm.G[storage.TaskRecord](s.db).Raw(`
		SELECT t.id, t.title, t.body, t.completed, t.meta, t.project_id, t.created_at, t.updated_at
		FROM tasks t
		JOIN board_tasks bt ON bt.task_id = t.id
		WHERE bt.board_id = ?;
    `, boardId).Find(ctx)
	if err != nil {
		return nil, err
	}

	taskIDs := make([]int, len(tasks))
	for i, t := range tasks {
		taskIDs[i] = t.ID
	}

	taskOwners, err := s.batchTaskOwnerRecord(ctx, taskIDs)
	if err != nil {
		return nil, err
	}
	taskAssignees, err := s.batchTaskAssigneeRecord(ctx, taskIDs)
	if err != nil {
		return nil, err
	}
	boardTasks, err := s.batchBoardTaskRecord(ctx, taskIDs)
	if err != nil {
		return nil, err
	}

	return s.toDomainModels(tasks, taskOwners, taskAssignees, boardTasks), nil
}

// ==============================
//            Helpers
// ==============================

func (s *gormTaskStorageRepository) toDomainModels(
	taskRecords []storage.TaskRecord,
	taskOwners []storage.TaskOwnerRecord,
	taskAssignees []storage.TaskAssigneeRecord,
	taskBoards []storage.BoardTaskRecord,
) []models.Task {
	tasks := make([]models.Task, len(taskRecords))
	for i, t := range taskRecords {
		task := models.Task{
			ID:           t.ID,
			Title:        t.Title,
			Body:         t.Body,
			Completed:    t.Completed,
			ProjectId:    t.ProjectId,
			Boards:       []models.TaskBoard{},
			AssigneesIDs: []int{},
			Meta:         t.Meta,
			CreatedAt:    t.CreatedAt,
			UpdatedAt:    t.UpdatedAt,
		}

		for _, o := range taskOwners {
			if o.TaskID == task.ID {
				task.OwnerId = o.UserID
				break
			}
		}

		for _, b := range taskBoards {
			if b.TaskID == task.ID {
				task.Boards = append(task.Boards, models.TaskBoard{
					BoardId:  b.BoardID,
					ColumnId: b.ColumnID,
					Position: b.Position,
				})
			}
		}

		for _, a := range taskAssignees {
			if a.TaskID == task.ID {
				task.AssigneesIDs = append(task.AssigneesIDs, a.UserID)
			}
		}

		tasks[i] = task
	}

	return tasks
}

func (s *gormTaskStorageRepository) batchTaskOwnerRecord(ctx context.Context, ids []int) ([]storage.TaskOwnerRecord, error) {
	return gorm.G[storage.TaskOwnerRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}

func (s *gormTaskStorageRepository) batchTaskAssigneeRecord(ctx context.Context, ids []int) ([]storage.TaskAssigneeRecord, error) {
	return gorm.G[storage.TaskAssigneeRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}

func (s *gormTaskStorageRepository) batchBoardTaskRecord(ctx context.Context, ids []int) ([]storage.BoardTaskRecord, error) {
	return gorm.G[storage.BoardTaskRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}

func (s *gormTaskStorageRepository) saveTaskRecord(ctx context.Context, tx *gorm.DB, task *models.Task) error {
	record := storage.TaskRecord{
		ID:        task.ID,
		Title:     task.Title,
		Body:      task.Body,
		Meta:      task.Meta,
		ProjectId: task.ProjectId,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if task.ID == 0 {
		if err := gorm.G[storage.TaskRecord](tx).Create(ctx, &record); err != nil {
			return err
		}
		task.ID = record.ID
		task.CreatedAt = record.CreatedAt
		task.UpdatedAt = record.UpdatedAt
		return nil
	}

	rows, err := gorm.G[storage.TaskRecord](tx).
		Where("id = ?", task.ID).
		Select("title", "body", "meta", "project_id", "completed", "updated_at").
		Updates(ctx, record)
	if err != nil {
		return err
	}
	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	task.UpdatedAt = record.UpdatedAt
	return nil
}

func (s *gormTaskStorageRepository) saveOwner(ctx context.Context, tx *gorm.DB, task *models.Task) error {
	if _, err := gorm.G[storage.TaskOwnerRecord](tx).Where("task_id = ?", task.ID).Delete(ctx); err != nil {
		return err
	}

	if task.OwnerId == 0 {
		return nil
	}

	return gorm.G[storage.TaskOwnerRecord](tx).Create(ctx, &storage.TaskOwnerRecord{
		TaskID: task.ID,
		UserID: task.OwnerId,
	})
}

func (s *gormTaskStorageRepository) saveAssignees(ctx context.Context, tx *gorm.DB, task *models.Task) error {
	if _, err := gorm.G[storage.TaskAssigneeRecord](tx).Where("task_id = ?", task.ID).Delete(ctx); err != nil {
		return err
	}

	seen := make(map[int]struct{}, len(task.AssigneesIDs))
	records := make([]storage.TaskAssigneeRecord, 0, len(task.AssigneesIDs))
	for _, userID := range task.AssigneesIDs {
		if userID == 0 {
			continue
		}
		if _, ok := seen[userID]; ok {
			continue
		}
		seen[userID] = struct{}{}
		records = append(records, storage.TaskAssigneeRecord{TaskID: task.ID, UserID: userID})
	}

	if len(records) == 0 {
		return nil
	}

	return gorm.G[storage.TaskAssigneeRecord](tx).CreateInBatches(ctx, &records, len(records))
}

func (s *gormTaskStorageRepository) saveBoards(ctx context.Context, tx *gorm.DB, task *models.Task) error {
	if _, err := gorm.G[storage.BoardTaskRecord](tx).Where("task_id = ?", task.ID).Delete(ctx); err != nil {
		return err
	}

	seen := make(map[int]struct{}, len(task.Boards))
	records := make([]storage.BoardTaskRecord, 0, len(task.Boards))
	for _, b := range task.Boards {
		if b.BoardId == 0 {
			continue
		}
		if _, ok := seen[b.BoardId]; ok {
			continue
		}
		seen[b.BoardId] = struct{}{}

		record := storage.BoardTaskRecord{
			BoardID:  b.BoardId,
			TaskID:   task.ID,
			Position: b.Position,
		}

		if b.ColumnId != 0 {
			columnID := b.ColumnId
			record.ColumnID = columnID
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		return nil
	}

	return gorm.G[storage.BoardTaskRecord](tx).CreateInBatches(ctx, &records, len(records))
}
