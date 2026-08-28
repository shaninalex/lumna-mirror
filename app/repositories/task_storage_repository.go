package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories/storage"
	"gorm.io/gorm"
)

type TaskStorageRepositoryFilterQuery struct {
	BoardId   *models.EntityId `json:"board_id" form:"board_id"`
	ProjectId *models.EntityId `json:"project_id" form:"project_id"`
}

type TaskStorageRepository interface {
	Filter(ctx context.Context, boardId models.EntityId) ([]models.Task, error)
	FindByID(ctx context.Context, id models.EntityId) (*models.Task, error)

	// Save creates or updates the entire cluster transactionally
	Save(ctx context.Context, task *models.Task) error

	// Explicit domain action optimized to bypass heavy loads when necessary
	AssignUser(ctx context.Context, taskID models.EntityId, userID models.IdentityID) error
	MovePosition(ctx context.Context, taskID, boardID, columnID models.EntityId, pos int64) error
	Complete(ctx context.Context, taskID models.EntityId) error
}

type gormTaskStorageRepository struct {
	db *gorm.DB
}

var _ TaskStorageRepository = (*gormTaskStorageRepository)(nil)

func NewGormTaskStorageRepository(db *gorm.DB) TaskStorageRepository {
	return &gormTaskStorageRepository{db: db}
}

// AssignUser implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) AssignUser(ctx context.Context, taskID models.EntityId, userID models.IdentityID) error {
	panic("unimplemented")
}

// Complete implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) Complete(ctx context.Context, taskID models.EntityId) error {
	panic("unimplemented")
}

// Filter implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) Filter(ctx context.Context, boardId models.EntityId) ([]models.Task, error) {
	panic("unimplemented")
}

// FindByID implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) FindByID(ctx context.Context, id models.EntityId) (*models.Task, error) {
	task, err := gorm.G[storage.TaskRecord](g.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := gorm.G[storage.TaskOwnerRecord](g.db).Where("task_id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	assignee, err := gorm.G[storage.TaskAssigneeRecord](g.db).Where("task_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	boards, err := gorm.G[storage.BoardTaskRecord](g.db).Where("task_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	t := g.toDomainModel(task, owner, assignee, boards)
	return &t, nil
}

// MovePosition implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) MovePosition(ctx context.Context, taskID, boardID, columnID models.EntityId, pos int64) error {
	panic("unimplemented")
}

// Save implements [TaskStorageRepository].
func (g *gormTaskStorageRepository) Save(ctx context.Context, task *models.Task) error {
	panic("unimplemented")
}

func (g *gormTaskStorageRepository) toDomainModel(
	taskRecord storage.TaskRecord,
	taskOwner storage.TaskOwnerRecord,
	taskAssignee []storage.TaskAssigneeRecord,
	taskBoards []storage.BoardTaskRecord,
) models.Task {
	return models.Task{}
}
