package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gorm.io/gorm"
)

type ColumnService struct {
	db               *gorm.DB
	columnRepository repositories.ColumnRepository
	boardRepository  repositories.BoardRepository
}

func NewColumnService(
	db *gorm.DB,
	boardRepository repositories.BoardRepository,
	columnRepository repositories.ColumnRepository,
) *ColumnService {
	return &ColumnService{
		db:               db,
		boardRepository:  boardRepository,
		columnRepository: columnRepository,
	}
}

func (s *ColumnService) Filter(ctx context.Context, boardId uint) []models.Column {
	return s.columnRepository.FilterByBoard(ctx, boardId)
}

func (s *ColumnService) Get(ctx context.Context, columnID uint) (*models.Column, error) {
	return s.columnRepository.GetByID(ctx, columnID)
}

// Replaced with update
//func (s *ColumnService) Reorder(ctx context.Context, columnID uint, order uint) error {
//	// TODO: get pointer of a struct and change it
//	return s.db.WithContext(ctx).Model(&models.Column{}).
//		Where("id = ?", columnID).
//		Update("order", order).Error
//}

type ColumnUpdate struct {
	BoardId uint   `json:"board_id"`
	Order   uint   `json:"order"`
	Title   string `json:"title"`
}

func (s *ColumnService) Create(ctx context.Context, payload ColumnUpdate) (*models.Column, error) {
	board, err := s.boardRepository.GetByID(ctx, payload.BoardId)
	if err != nil {
		return nil, err
	}

	column := models.Column{
		BoardID:   board.ID,
		Order:     payload.Order,
		Title:     payload.Title,
		ProjectID: board.ProjectID,
	}

	return s.columnRepository.Create(ctx, column)
}

func (s *ColumnService) Update(ctx context.Context, column *models.Column) (*models.Column, error) {
	return s.columnRepository.Update(ctx, column)
}

func (s *ColumnService) Delete(ctx context.Context, id uint) error {
	return s.columnRepository.Delete(ctx, id)
}
