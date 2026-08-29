package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories/storage"
	"gorm.io/gorm"
)

type TaskStorageRepositoryFilterQuery struct {
	BoardId   *int64 `json:"board_id" form:"board_id"`
	ProjectId *int64 `json:"project_id" form:"project_id"`
}

type TaskStorageRepository interface {
	Filter(ctx context.Context, boardId int64) ([]models.Task, error)
	FindByID(ctx context.Context, id int64) (*models.Task, error)

	// Save creates or updates the entire cluster transactionally
	Save(ctx context.Context, task *models.Task) error

	// Explicit domain action optimized to bypass heavy loads when necessary
	AssignUser(ctx context.Context, taskID int64, userID models.IdentityID) error
	MovePosition(ctx context.Context, taskID, boardID, columnID int64, pos int64) error
	Complete(ctx context.Context, taskID int64) error
}

type gormTaskStorageRepository struct {
	db *gorm.DB
}

var _ TaskStorageRepository = (*gormTaskStorageRepository)(nil)

func NewGormTaskStorageRepository(db *gorm.DB) TaskStorageRepository {
	return &gormTaskStorageRepository{db: db}
}

// AssignUser implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) AssignUser(ctx context.Context, taskID int64, userID models.IdentityID) error {
	// TODO: create storage.TaskAssigneeRecord
	panic("unimplemented")
}

// Complete implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Complete(ctx context.Context, taskID int64) error {
	// TODO: update storage.TaskRecord
	panic("unimplemented")
}

// FindByID implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) FindByID(ctx context.Context, id int64) (*models.Task, error) {
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
func (s *gormTaskStorageRepository) MovePosition(ctx context.Context, taskID, boardID, columnID int64, pos int64) error {
	// TODO: update storage.TaskBoardRecord
	panic("unimplemented")
}

// Save implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Save(ctx context.Context, task *models.Task) error {
	// TODO: get data from model and update records.
	// Important: something might need delete and create again, depends of what we are doing - save or update
	panic("unimplemented")
}

// Filter implements [TaskStorageRepository].
func (s *gormTaskStorageRepository) Filter(ctx context.Context, boardId int64) ([]models.Task, error) {
	tasks, err := gorm.G[storage.TaskRecord](s.db).Raw(`
		SELECT t.id, t.title, t.body, t.completed, t.meta, t.project_id, t.created_at, t.updated_at
		FROM tasks t
		JOIN board_tasks bt ON bt.task_id = t.id
		WHERE bt.board_id = ?;
    `, boardId).Find(ctx)
	if err != nil {
		return nil, err
	}

	taskIDs := make([]int64, len(tasks))
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

func (s *gormTaskStorageRepository) toDomainModels(
	taskRecords []storage.TaskRecord,
	taskOwners []storage.TaskOwnerRecord,
	taskAssignees []storage.TaskAssigneeRecord,
	taskBoards []storage.BoardTaskRecord,
) []models.Task {
	tasks := make([]models.Task, len(taskRecords))

	for _, t := range taskRecords {
		task := models.Task{
			ID:           t.ID,
			Title:        t.Title,
			Body:         t.Body,
			Completed:    t.Completed,
			ProjectId:    t.ProjectId,
			Boards:       []models.TaskBoard{},
			AssigneesIDs: []int64{},
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
					ColumnId: *b.ColumnID,
					Position: b.Position,
				})
			}
		}

		for _, a := range taskAssignees {
			if a.TaskID == task.ID {
				task.AssigneesIDs = append(task.AssigneesIDs, a.UserID)
			}
		}
	}

	return tasks
}

func (s *gormTaskStorageRepository) batchTaskOwnerRecord(ctx context.Context, ids []int64) ([]storage.TaskOwnerRecord, error) {
	return gorm.G[storage.TaskOwnerRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}

func (s *gormTaskStorageRepository) batchTaskAssigneeRecord(ctx context.Context, ids []int64) ([]storage.TaskAssigneeRecord, error) {
	return gorm.G[storage.TaskAssigneeRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}

func (s *gormTaskStorageRepository) batchBoardTaskRecord(ctx context.Context, ids []int64) ([]storage.BoardTaskRecord, error) {
	return gorm.G[storage.BoardTaskRecord](s.db).Where("task_id IN ?", ids).Find(ctx)
}
