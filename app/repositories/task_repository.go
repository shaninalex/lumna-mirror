package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	List(ctx context.Context, query map[string]any) ([]*models.Task, error)
	GetByID(ctx context.Context, taskID uint) (*models.Task, error)
	Reorder(ctx context.Context, taskID uint, columnID uint, order uint) error
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	UpdateFields(ctx context.Context, taskID uint, updates map[string]any) error
}

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) TaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) List(ctx context.Context, query map[string]any) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.Where(query).Find(&tasks).Error; err != nil {
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

func (r *GormTaskRepository) Reorder(ctx context.Context, taskID uint, columnID uint, order uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]any{
			"column_id": columnID,
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
