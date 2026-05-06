package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type BoardService struct {
	boardRepository  repositories.BoardRepository
	taskRepository   repositories.TaskRepository
	columnRepository repositories.ColumnRepository
}

func NewBoardService(
	boardRepository repositories.BoardRepository,
	taskRepository repositories.TaskRepository,
	columnRepository repositories.ColumnRepository,
) *BoardService {
	return &BoardService{
		boardRepository:  boardRepository,
		taskRepository:   taskRepository,
		columnRepository: columnRepository,
	}
}

func (s *BoardService) Create(ctx context.Context, projectID uint, title string) (*models.Board, error) {
	board := models.Board{
		ProjectID: projectID,
		Title:     title,
	}
	return s.boardRepository.Create(ctx, &board)
}

func (s *BoardService) Delete(ctx context.Context, boardID uint) error {
	return s.boardRepository.Delete(ctx, boardID)
}

func (s *BoardService) Update(ctx context.Context, board *models.Board) error {
	_, err := s.boardRepository.Update(ctx, board)
	return err
}

func (s *BoardService) Get(ctx context.Context, boardID uint) (*models.Board, error) {
	return s.boardRepository.GetByID(ctx, boardID)
}
func (s *BoardService) List(ctx context.Context, projectID uint) ([]*models.Board, error) {
	return s.boardRepository.ListByProjectId(ctx, projectID)
}

type KTask struct {
	Id    uint `json:"id"`
	Order uint `json:"order"`
}

type KColumn struct {
	Id    uint    `json:"id"`
	Order *uint   `json:"order"`
	Tasks []KTask `json:"tasks"`
}

type KanbanBoardChangeOrderPayload struct {
	MoveType string                     `json:"moveType"`
	ColumnId *uint                      `json:"columnId"`
	Tasks    []KTask                    `json:"tasks"`
	Columns  []KColumn                  `json:"columns"`
	Activity *logger.ActivityLogPayload `json:"activity"`
}

func (s *BoardService) Reorder(ctx context.Context, payload *KanbanBoardChangeOrderPayload) error {
	if payload.MoveType == "task" {
		// move in a single column
		if payload.ColumnId != nil {
			for _, task := range payload.Tasks {
				if err := s.taskRepository.UpdateFields(ctx, task.Id, map[string]any{
					"order": task.Order,
				}); err != nil {
					return err
				}
			}
		} else {
			// move task between columns
			for _, column := range payload.Columns {
				for _, task := range column.Tasks {
					if err := s.taskRepository.UpdateFields(ctx, task.Id, map[string]any{
						"order":     task.Order,
						"column_id": column.Id,
					}); err != nil {
						return err
					}
				}
			}
		}
	} else {
		// move column between each other
		for _, column := range payload.Columns {
			if err := s.columnRepository.UpdateFields(ctx, column.Id, map[string]any{"order": column.Order}); err != nil {
				return err
			}
		}
	}

	return nil
}
