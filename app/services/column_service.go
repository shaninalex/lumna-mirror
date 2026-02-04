package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type ColumnService struct {
}

func NewColumnService() *ColumnService {
	return &ColumnService{}
}

func (s *ColumnService) Filter(ctx context.Context, boardId uuid.UUID) []models.Column {
	var columns []models.Column
	if result := db.GetDB(ctx).Preload("Tasks").Where("board_id = ?", boardId).Find(&columns); result.Error != nil {
		return []models.Column{}
	}
	return columns
}

func (s *ColumnService) Get(ctx context.Context, boardListID uuid.UUID) (*models.Column, error) {
	var column models.Column
	if result := db.GetDB(ctx).Where("id = ?", boardListID).First(&column); result.Error != nil {
		return nil, result.Error
	}
	return &column, nil
}

func (s *ColumnService) Reorder(ctx context.Context, listID uuid.UUID, order uint) error {
	// TODO: get pointer of a struct and change it
	return db.GetDB(ctx).Model(&models.Column{}).
		Where("id = ?", listID).
		Update("order", order).Error
}

type NewColumn struct {
	BoardId uuid.UUID `json:"board_id"`
	Order   uint      `json:"order"`
	Title   string    `json:"title"`
}

func (s *ColumnService) Create(ctx context.Context, payload NewColumn) (*models.Column, error) {
	column := models.Column{
		BoardID: payload.BoardId,
		Order:   payload.Order,
		Title:   payload.Title,
	}
	if result := db.GetDB(ctx).Create(&column); result.Error != nil {
		return nil, result.Error
	}
	return &column, nil
}

func (s *ColumnService) Delete(ctx context.Context, id uuid.UUID) error {
	if result := db.GetDB(ctx).Where("id = ?", id).Delete(&models.Column{}); result.Error != nil {
		return result.Error
	}
	return nil
}
