package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ColumnService struct {
	db *gorm.DB
}

func NewColumnService(db *gorm.DB) *ColumnService {
	return &ColumnService{db: db}
}

func (s *ColumnService) Filter(ctx context.Context, boardId uint) []models.Column {
	var columns []models.Column
	if result := s.db.WithContext(ctx).Preload("Tasks").Where("board_id = ?", boardId).Find(&columns); result.Error != nil {
		return []models.Column{}
	}
	return columns
}

func (s *ColumnService) Get(ctx context.Context, boardListID uint) (*models.Column, error) {
	var column models.Column
	if result := s.db.WithContext(ctx).Where("id = ?", boardListID).First(&column); result.Error != nil {
		return nil, result.Error
	}
	return &column, nil
}

func (s *ColumnService) Reorder(ctx context.Context, listID uint, order uint) error {
	// TODO: get pointer of a struct and change it
	return s.db.WithContext(ctx).Model(&models.Column{}).
		Where("id = ?", listID).
		Update("order", order).Error
}

type ColumnUpdate struct {
	BoardId uint   `json:"board_id"`
	Order   uint   `json:"order"`
	Title   string `json:"title"`
}

func (s *ColumnService) Create(ctx context.Context, payload ColumnUpdate) (*models.Column, error) {
	board := &models.Board{}
	result := s.db.WithContext(ctx).Where("id = ?", payload.BoardId).First(board)
	if result.Error != nil {
		return nil, result.Error
	}

	column := models.Column{
		BoardID:   board.ID,
		Order:     payload.Order,
		Title:     payload.Title,
		ProjectID: board.ProjectID,
	}
	if result := s.db.WithContext(ctx).Create(&column); result.Error != nil {
		return nil, result.Error
	}
	return &column, nil
}

func (s *ColumnService) Update(ctx context.Context, column *models.Column) (*models.Column, error) {
	if result := s.db.WithContext(ctx).Save(&column); result.Error != nil {
		return nil, result.Error
	}
	return column, nil
}

func (s *ColumnService) Delete(ctx context.Context, id uint) error {
	if result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Column{}); result.Error != nil {
		return result.Error
	}
	return nil
}
