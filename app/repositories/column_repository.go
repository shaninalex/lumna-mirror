package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ColumnRepository interface {
	Create(ctx context.Context, column models.Column) (*models.Column, error)
	Update(ctx context.Context, column *models.Column) (*models.Column, error)
	UpdateFields(ctx context.Context, columnID uint, updates map[string]any) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, columnID uint) (*models.Column, error)
	FilterByBoard(ctx context.Context, boardId uint) []models.Column
}

type GormColumnRepository struct {
	db *gorm.DB
}

func NewGormColumnRepository(db *gorm.DB) ColumnRepository {
	return &GormColumnRepository{
		db: db,
	}
}

func (r *GormColumnRepository) GetByID(ctx context.Context, id uint) (*models.Column, error) {
	var column models.Column

	err := r.db.WithContext(ctx).
		First(&column, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &column, nil
}

func (r *GormColumnRepository) FilterByBoard(ctx context.Context, boardId uint) []models.Column {
	var columns []models.Column
	if result := r.db.WithContext(ctx).
		//Preload("Tasks").
		Where("board_id = ?", boardId).
		Find(&columns); result.Error != nil {
		return []models.Column{}
	}
	return columns
}

func (r *GormColumnRepository) Create(ctx context.Context, column models.Column) (*models.Column, error) {
	if result := r.db.WithContext(ctx).Create(&column); result.Error != nil {
		return nil, result.Error
	}
	return &column, nil
}

func (r *GormColumnRepository) Update(ctx context.Context, column *models.Column) (*models.Column, error) {
	if result := r.db.WithContext(ctx).Save(&column); result.Error != nil {
		return nil, result.Error
	}
	return column, nil
}

func (r *GormColumnRepository) Delete(ctx context.Context, id uint) error {
	if result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Column{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormColumnRepository) UpdateFields(ctx context.Context, columnID uint, updates map[string]any) error {
	result := r.db.WithContext(ctx).Model(&models.Column{}).
		Where("id = ?", columnID).
		Updates(updates)
	return result.Error
}
