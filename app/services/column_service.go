package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type BoardCreatePayload struct {
	Title   string `json:"title"`
	Order   int64  `json:"order"`
	BoardId int64  `json:"board_id"`
}

type ColumnService interface {
	Filter(ctx context.Context, boardId uint) []models.Column
	Get(ctx context.Context, columnId uint) (*models.Column, error)
	Save(ctx context.Context, column *models.Column) error
	Delete(ctx context.Context, id uint) error
}

type columnService struct {
	statusRepository repositories.ColumnRepository
	listRepository   repositories.BoardRepository
}

func NewStatusService(
	listRepository repositories.BoardRepository,
	statusRepository repositories.ColumnRepository,
) ColumnService {
	return &columnService{
		listRepository:   listRepository,
		statusRepository: statusRepository,
	}
}

func (s *columnService) Filter(ctx context.Context, boardId uint) []models.Column {
	return s.statusRepository.FilterByBoard(ctx, boardId)
}

func (s *columnService) Get(ctx context.Context, columnId uint) (*models.Column, error) {
	return s.statusRepository.GetByID(ctx, columnId)
}

func (s *columnService) Save(ctx context.Context, column *models.Column) error {
	return s.statusRepository.Save(ctx, column)
}

func (s *columnService) Delete(ctx context.Context, id uint) error {
	return s.statusRepository.Delete(ctx, id)
}
