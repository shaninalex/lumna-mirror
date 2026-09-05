package services

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

var (
	EmptyColumnListError error = errors.New("empty column list")
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
	Reorder(ctx context.Context, ids []int64) error
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

// Reorder implements [ColumnService].
func (s *columnService) Reorder(ctx context.Context, columnIds []int64) error {
	if len(columnIds) == 0 {
		return EmptyColumnListError
	}

	for i, columnId := range columnIds {
		column, err := s.Get(ctx, uint(columnId))
		if err != nil {
			return err
		}
		column.Position = int64(i)
		if err := s.Save(ctx, column); err != nil {
			return err
		}
	}

	return nil
}
