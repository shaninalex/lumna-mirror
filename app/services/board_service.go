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

func (s *BoardService) GetBoard(ctx context.Context, boardId uint) (*models.Board, error) {
	return s.boardRepository.Get(ctx, boardId)
}

func (s *BoardService) ProjectBoards(ctx context.Context, projectID uint) ([]*models.Board, error) {
	return s.boardRepository.List(ctx, db.Eq("project_id", projectID))
}

func (s *BoardService) Update(ctx context.Context, id uint, opts db.SetExpr) error {
	if err := s.boardRepository.Update(ctx, id, opts); err != nil {
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
	boardList, err := s.boardListRepository.List(ctx, db.Eq("board_id", id))
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

func (s *BoardService) ListUpdate(ctx context.Context, id uint, opts db.SetExpr) error {
	return s.boardListRepository.Update(ctx, id, opts)
}

func (s *BoardService) ListDelete(ctx context.Context, id uint) error {
	return s.boardListRepository.Delete(ctx, id)
}

func (s *BoardService) ListGet(ctx context.Context, id uint) (*models.BoardList, error) {
	return s.boardListRepository.Get(ctx, id)
}

type Tasks struct {
	ID    uint `json:"id"`
	Order uint `json:"order"`
}

type List struct {
	ID    uint    `json:"id"`
	Order *uint   `json:"order"`
	Tasks []Tasks `json:"tasks"`
}

type ChangeOrderPayload struct {
	Lists []List `json:"lists"`
}

func (s *BoardService) ChangeOrder(ctx context.Context, boardId uint, payload *ChangeOrderPayload) error {
	for _, list := range payload.Lists {
		// TODO: this is bad. Create single query.
		s.boardListRepository.Update(ctx, list.ID, db.Set(
			db.Field("list_order", list.Order),
		))
	}
	return nil
}
