package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ColumnService interface {
	Filter(ctx context.Context, boardId uint) []models.Column
	Get(ctx context.Context, columnId uint) (*models.Column, error)
	Create(ctx context.Context, payload BoardUpdate) (*models.Column, error)
	Update(ctx context.Context, column *models.Column) (*models.Column, error)
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

type BoardUpdate struct {
	Title   string `json:"title"`
	Order   uint   `json:"order"`
	BoardId uint   `json:"board_id"`
}

func (s *columnService) Create(ctx context.Context, payload BoardUpdate) (*models.Column, error) {
	status := models.Column{
		Title:   payload.Title,
		BoardId: payload.BoardId,
	}
	return s.statusRepository.Create(ctx, status)
}

func (s *columnService) Update(ctx context.Context, status *models.Column) (*models.Column, error) {
	return s.statusRepository.Update(ctx, status)
}

func (s *columnService) Delete(ctx context.Context, id uint) error {
	return s.statusRepository.Delete(ctx, id)
}
