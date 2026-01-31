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

func (s *ColumnService) ReorderList(ctx context.Context, listID uuid.UUID, order uint) error {
	database := db.GetDB(ctx)
	return database.Model(&models.Column{}).
		Where("id = ?", listID).
		Update("order", order).Error
}

func (s *ColumnService) BoardListGet(ctx context.Context, boardListID uuid.UUID) (*models.Column, error) {
	var boardList models.Column
	if result := db.GetDB(ctx).Where("id = ?", boardListID).First(&boardList); result.Error != nil {
		return nil, result.Error
	}
	return &boardList, nil
}
