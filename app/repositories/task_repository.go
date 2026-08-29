package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type TaskListQuery struct {
	ProjectID *uint
	BoardId   *uint
	Code      *string
	QueryArgs []interface{}
	OrderBy   *string
	Limit     *uint
}

// TaskRepository
//
// Deprecated: the repository interface is not flexible enough for new database schema.
type TaskRepository interface {
	Repository
	List(ctx context.Context, query TaskListQuery) ([]*models.Task, error)
	GetByID(ctx context.Context, taskID uint) (*models.Task, error)
	Reorder(ctx context.Context, taskID uint, statusID uint, order uint) error
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	UpdateFields(ctx context.Context, taskID uint, updates map[string]any) error
	CreateTaskBoard(ctx context.Context, boardTask models.BoardTask) error
	GetTaskBoards(ctx context.Context, boardId uint) ([]models.BoardTask, error)
	GetTasksByBoardId(ctx context.Context, boardId uint) ([]models.Task, error)
	AddTaskToBoard(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error)
}

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) TaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *GormTaskRepository) List(ctx context.Context, query TaskListQuery) ([]*models.Task, error) {
	var tasks []*models.Task
	db := r.db.WithContext(ctx)
	if query.ProjectID != nil {
		db = db.Where("project_id = ?", *query.ProjectID)
	}

	if query.OrderBy != nil {
		db = db.Order(*query.OrderBy)
	}

	if query.Limit != nil {
		db = db.Limit(int(*query.Limit))
	}
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *GormTaskRepository) GetByID(ctx context.Context, taskID uint) (*models.Task, error) {
	var task models.Task

	if err := r.db.WithContext(ctx).
		Where("id = ?", taskID).
		First(&task).
		Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *GormTaskRepository) Reorder(ctx context.Context, taskID uint, statusID uint, order uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]any{
			"status_id": statusID,
			"order":     order,
		}).Error
}

func (r *GormTaskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).
		Create(task).
		Error
}

func (r *GormTaskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).
		Save(task).
		Error
}

func (r *GormTaskRepository) UpdateFields(ctx context.Context, taskID uint, updates map[string]any) error {
	result := r.db.WithContext(ctx).Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(updates)
	return result.Error
}

// CreateTaskBoard implements [TaskRepository].
func (r *GormTaskRepository) CreateTaskBoard(ctx context.Context, boardTask models.BoardTask) error {
	return gorm.G[models.BoardTask](r.db).Create(ctx, &boardTask)
}

// GetTaskBoards implements [TaskRepository].
func (r *GormTaskRepository) GetTaskBoards(ctx context.Context, boardId uint) ([]models.BoardTask, error) {
	return gorm.G[models.BoardTask](r.db).Raw(`
        SELECT board_id, task_id, column_id, position
		FROM board_tasks
        WHERE board_id = ?
    `, boardId).Find(ctx)
}

// GetTasksByBoardId implements [TaskRepository].
func (r *GormTaskRepository) GetTasksByBoardId(ctx context.Context, boardId uint) ([]models.Task, error) {
	return gorm.G[models.Task](r.db).Raw(`
		SELECT id, title, body, completed, meta, created_at, updated_at
		FROM tasks t
		JOIN board_tasks bt ON bt.task_id = t.id
		JOIN boards b ON b.id = bt.board_id
		WHERE bt.board_id = 1;
    `, boardId).Find(ctx)
}

// AddTaskToBoard implements [TaskRepository].
func (r *GormTaskRepository) AddTaskToBoard(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error) {
	task := models.Task{
		Title: payload.Title,
		Body:  payload.Body,
	}

	// TODO: <= transaction start
	if err := r.Create(ctx, &task); err != nil {
		return nil, err
	}

	if err := r.CreateTaskBoard(ctx, models.BoardTask{
		TaskId:   uint(task.ID),
		Position: uint(0),
		BoardId:  uint(payload.BoardId),
		ColumnId: uint(payload.ColumnId),
	}); err != nil {
		return nil, err
	}
	// TODO: <= transaction commit

	return &task, nil
}
