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
	Order   int    `json:"order"`
	BoardId int    `json:"board_id"`
}

type ColumnService interface {
	Filter(ctx context.Context, boardId int) []models.Column
	Get(ctx context.Context, columnId int) (*models.Column, error)
	Save(ctx context.Context, column *models.Column) error
	Delete(ctx context.Context, id int) error
	Reorder(ctx context.Context, ids []int) error
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

func (s *columnService) Filter(ctx context.Context, boardId int) []models.Column {
	return s.statusRepository.FilterByBoard(ctx, boardId)
}

func (s *columnService) Get(ctx context.Context, columnId int) (*models.Column, error) {
	return s.statusRepository.GetByID(ctx, columnId)
}

func (s *columnService) Save(ctx context.Context, column *models.Column) error {
	return s.statusRepository.Save(ctx, column)
}

func (s *columnService) Delete(ctx context.Context, id int) error {
	return s.statusRepository.Delete(ctx, id)
}

// Reorder implements [ColumnService].
func (s *columnService) Reorder(ctx context.Context, columnIds []int) error {
	if len(columnIds) == 0 {
		return EmptyColumnListError
	}

	for i, columnId := range columnIds {
		column, err := s.Get(ctx, columnId)
		if err != nil {
			return err
		}
		column.Position = i
		if err := s.Save(ctx, column); err != nil {
			return err
		}
	}

	return nil
}
