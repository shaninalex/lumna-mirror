package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type BoardService struct {
	boardRepository     *repositories.BoardRepository
	boardListRepository *repositories.BoardListRepository
}

func NewBoardService() *BoardService {
	return &BoardService{
		boardRepository:     repositories.NewBoardRepository(),
		boardListRepository: repositories.NewBoardListRepository(),
	}
}

func (s *BoardService) ProjectBoards(ctx context.Context, projectID uint) ([]*models.Board, error) {
	return s.boardRepository.List(ctx, db.Option{Key: "project_id", Value: projectID})
}

func (s *BoardService) Update(ctx context.Context, id uint, opts ...db.Option) error {
	if err := s.boardRepository.Update(ctx, id, opts...); err != nil {
		return err
	}
	return nil
}

func (s *BoardService) Delete(ctx context.Context, id uint) error {
	return s.boardRepository.Delete(ctx, id)
}

func (s *BoardService) Create(ctx context.Context, entry *models.Board) (*models.Board, error) {
	if err := s.boardRepository.Create(ctx, entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (s *BoardService) Lists(ctx context.Context, id uint) ([]*models.BoardList, error) {
	boardList, err := s.boardListRepository.List(ctx, db.Option{Key: "board_id", Value: id})
	if err != nil {
		return nil, err
	}
	return boardList, nil
}

func (s *BoardService) ListCreate(ctx context.Context, entry *models.BoardList) (*models.BoardList, error) {
	if err := s.boardListRepository.Create(ctx, entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (s *BoardService) ListUpdate(ctx context.Context, id uint, opts ...db.Option) error {
	return s.boardListRepository.Update(ctx, id, opts...)
}

func (s *BoardService) ListDelete(ctx context.Context, id uint) error {
	return s.boardListRepository.Delete(ctx, id)
}
