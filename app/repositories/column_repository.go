package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ColumnRepository interface {
	Repository
	Save(ctx context.Context, status *models.Column) error
	UpdateFields(ctx context.Context, statusID int, updates map[string]any) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, statusID int) (*models.Column, error)
	FilterByBoard(ctx context.Context, listId int) []models.Column
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

func (r *GormColumnRepository) GetByID(ctx context.Context, id int) (*models.Column, error) {
	var status models.Column

	err := r.db.WithContext(ctx).
		First(&status, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *GormColumnRepository) FilterByBoard(ctx context.Context, listId int) []models.Column {
	var statuss []models.Column
	if result := r.db.WithContext(ctx).
		Where("board_id = ?", listId).
		Find(&statuss); result.Error != nil {
		return []models.Column{}
	}
	return statuss
}

func (r *GormColumnRepository) Save(ctx context.Context, column *models.Column) error {
	return r.db.WithContext(ctx).Save(&column).Error
}

func (r *GormColumnRepository) Delete(ctx context.Context, id int) error {
	if result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Column{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormColumnRepository) UpdateFields(ctx context.Context, statusID int, updates map[string]any) error {
	result := r.db.WithContext(ctx).Model(&models.Column{}).
		Where("id = ?", statusID).
		Updates(updates)
	return result.Error
}
