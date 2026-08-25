package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type BoardService interface {
	Create(ctx context.Context, projectID uint, title string) (*models.Board, error)
	Delete(ctx context.Context, listID uint) error
	Update(ctx context.Context, list *models.Board) error
	Get(ctx context.Context, listID uint) (*models.Board, error)
	List(ctx context.Context, projectID uint) ([]*models.Board, error)
	Reorder(ctx context.Context, payload *KanbanListChangeOrderPayload) error
}

type boardService struct {
	listRepository   repositories.BoardRepository
	taskRepository   repositories.TaskRepository
	statusRepository repositories.ColumnRepository
}

var _ BoardService = (*boardService)(nil)

func NewListService(
	listRepository repositories.BoardRepository,
	taskRepository repositories.TaskRepository,
	statusRepository repositories.ColumnRepository,
) BoardService {
	return &boardService{
		listRepository:   listRepository,
		taskRepository:   taskRepository,
		statusRepository: statusRepository,
	}
}

func (s *boardService) Create(ctx context.Context, projectID uint, title string) (*models.Board, error) {
	list := models.Board{
		ProjectID: projectID,
		Title:     title,
	}
	return s.listRepository.Create(ctx, &list)
}

func (s *boardService) Delete(ctx context.Context, listID uint) error {
	return s.listRepository.Delete(ctx, listID)
}

func (s *boardService) Update(ctx context.Context, list *models.Board) error {
	_, err := s.listRepository.Update(ctx, list)
	return err
}

func (s *boardService) Get(ctx context.Context, listID uint) (*models.Board, error) {
	return s.listRepository.GetByID(ctx, listID)
}
func (s *boardService) List(ctx context.Context, projectID uint) ([]*models.Board, error) {
	return s.listRepository.ListByProjectId(ctx, projectID)
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

type KanbanListChangeOrderPayload struct {
	MoveType string    `json:"move_type"`
	StatusId *uint     `json:"status_id"`
	Tasks    []KTask   `json:"tasks"`
	Statuses []KColumn `json:"statuses"`
}

func (s *boardService) Reorder(ctx context.Context, payload *KanbanListChangeOrderPayload) error {
	if payload.MoveType == "task" {
		// move in a single status
		if payload.StatusId != nil {
			for _, task := range payload.Tasks {
				if err := s.taskRepository.UpdateFields(ctx, task.Id, map[string]any{
					"order": task.Order,
				}); err != nil {
					return err
				}
			}
		} else {
			// move task between statuss
			for _, status := range payload.Statuses {
				for _, task := range status.Tasks {
					if err := s.taskRepository.UpdateFields(ctx, task.Id, map[string]any{
						"order":     task.Order,
						"status_id": status.Id,
					}); err != nil {
						return err
					}
				}
			}
		}
	} else {
		// move status between each other
		for _, status := range payload.Statuses {
			if err := s.statusRepository.UpdateFields(ctx, status.Id, map[string]any{"order": status.Order}); err != nil {
				return err
			}
		}
	}

	return nil
}
