package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type StatusRepository interface {
	Repository
	Create(ctx context.Context, status models.Status) (*models.Status, error)
	Update(ctx context.Context, status *models.Status) (*models.Status, error)
	UpdateFields(ctx context.Context, statusID uint, updates map[string]any) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, statusID uint) (*models.Status, error)
	FilterByList(ctx context.Context, listId uint) []models.Status
}

type GormStatusRepository struct {
	db *gorm.DB
}

func NewGormStatusRepository(db *gorm.DB) StatusRepository {
	return &GormStatusRepository{
		db: db,
	}
}

func (r *GormStatusRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *GormStatusRepository) GetByID(ctx context.Context, id uint) (*models.Status, error) {
	var status models.Status

	err := r.db.WithContext(ctx).
		First(&status, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *GormStatusRepository) FilterByList(ctx context.Context, listId uint) []models.Status {
	var statuss []models.Status
	if result := r.db.WithContext(ctx).
		//Preload("Tasks").
		Where("list_id = ?", listId).
		Find(&statuss); result.Error != nil {
		return []models.Status{}
	}
	return statuss
}

func (r *GormStatusRepository) Create(ctx context.Context, status models.Status) (*models.Status, error) {
	if result := r.db.WithContext(ctx).Create(&status); result.Error != nil {
		return nil, result.Error
	}
	return &status, nil
}

func (r *GormStatusRepository) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	if result := r.db.WithContext(ctx).Save(&status); result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}

func (r *GormStatusRepository) Delete(ctx context.Context, id uint) error {
	if result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Status{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormStatusRepository) UpdateFields(ctx context.Context, statusID uint, updates map[string]any) error {
	result := r.db.WithContext(ctx).Model(&models.Status{}).
		Where("id = ?", statusID).
		Updates(updates)
	return result.Error
}
