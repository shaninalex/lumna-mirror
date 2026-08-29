package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ColumnRepository interface {
	Repository
	Create(ctx context.Context, status *models.Column) error
	Update(ctx context.Context, status *models.Column) (*models.Column, error)
	UpdateFields(ctx context.Context, statusID uint, updates map[string]any) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, statusID uint) (*models.Column, error)
	FilterByBoard(ctx context.Context, listId uint) []models.Column
}

type GormColumnRepository struct {
	db *gorm.DB
}

func NewGormColumnRepository(db *gorm.DB) ColumnRepository {
	return &GormColumnRepository{
		db: db,
	}
}

func (r *GormColumnRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *GormColumnRepository) GetByID(ctx context.Context, id uint) (*models.Column, error) {
	var status models.Column

	err := r.db.WithContext(ctx).
		First(&status, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *GormColumnRepository) FilterByBoard(ctx context.Context, listId uint) []models.Column {
	var statuss []models.Column
	if result := r.db.WithContext(ctx).
		//Preload("Tasks").
		Where("board_id = ?", listId).
		Find(&statuss); result.Error != nil {
		return []models.Column{}
	}
	return statuss
}

func (r *GormColumnRepository) Create(ctx context.Context, column *models.Column) error {
	if err := r.db.WithContext(ctx).Create(&column).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormColumnRepository) Update(ctx context.Context, status *models.Column) (*models.Column, error) {
	if result := r.db.WithContext(ctx).Save(&status); result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}

func (r *GormColumnRepository) Delete(ctx context.Context, id uint) error {
	if result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Column{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormColumnRepository) UpdateFields(ctx context.Context, statusID uint, updates map[string]any) error {
	result := r.db.WithContext(ctx).Model(&models.Column{}).
		Where("id = ?", statusID).
		Updates(updates)
	return result.Error
}
