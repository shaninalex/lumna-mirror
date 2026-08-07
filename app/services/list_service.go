package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ListService interface {
	Create(ctx context.Context, projectID uint, title string) (*models.List, error)
	Delete(ctx context.Context, listID uint) error
	Update(ctx context.Context, list *models.List) error
	Get(ctx context.Context, listID uint) (*models.List, error)
	List(ctx context.Context, projectID uint) ([]*models.List, error)
	Reorder(ctx context.Context, payload *KanbanListChangeOrderPayload) error
}

type listService struct {
	listRepository   repositories.ListRepository
	taskRepository   repositories.TaskRepository
	statusRepository repositories.StatusRepository
}

var _ ListService = (*listService)(nil)

func NewListService(
	listRepository repositories.ListRepository,
	taskRepository repositories.TaskRepository,
	statusRepository repositories.StatusRepository,
) ListService {
	return &listService{
		listRepository:   listRepository,
		taskRepository:   taskRepository,
		statusRepository: statusRepository,
	}
}

func (s *listService) Create(ctx context.Context, projectID uint, title string) (*models.List, error) {
	list := models.List{
		ProjectID: projectID,
		Title:     title,
	}
	return s.listRepository.Create(ctx, &list)
}

func (s *listService) Delete(ctx context.Context, listID uint) error {
	return s.listRepository.Delete(ctx, listID)
}

func (s *listService) Update(ctx context.Context, list *models.List) error {
	_, err := s.listRepository.Update(ctx, list)
	return err
}

func (s *listService) Get(ctx context.Context, listID uint) (*models.List, error) {
	return s.listRepository.GetByID(ctx, listID)
}
func (s *listService) List(ctx context.Context, projectID uint) ([]*models.List, error) {
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

func (s *listService) Reorder(ctx context.Context, payload *KanbanListChangeOrderPayload) error {
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
