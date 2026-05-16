package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type TaskService struct {
	repository       repositories.TaskRepository
	statusRepository repositories.StatusRepository
}

func NewTaskService(
	repository repositories.TaskRepository,
	statusRepository repositories.StatusRepository,
) *TaskService {
	return &TaskService{
		repository:       repository,
		statusRepository: statusRepository,
	}
}

func (s *TaskService) List(ctx context.Context, query map[string]any) ([]*models.Task, error) {
	return s.repository.List(ctx, query)
}

func (s *TaskService) GetTask(ctx context.Context, taskID uint) (*models.Task, error) {
	return s.repository.GetByID(ctx, taskID)
}

func (s *TaskService) ReorderTask(ctx context.Context, taskID uint, listListID uint, order uint) error {
	return s.repository.Reorder(ctx, taskID, listListID, order)
}

// TaskPayload - used to create/partial update task
// TODO: add validators
type TaskPayload struct {
	Title     string `json:"title"`
	Order     uint   `json:"order"`
	ProjectID uint   `json:"project_id"`
	ColumnID  uint   `json:"status_id"`
}

func (s *TaskService) CreateTask(ctx context.Context, payload *TaskPayload) (*models.Task, error) {
	status, err := s.statusRepository.GetByID(ctx, payload.ColumnID)
	if err != nil {
		return nil, err
	}
	task := models.Task{
		Title:     payload.Title,
		Order:     payload.Order,
		StatusID:  status.ID,
		ProjectID: status.ProjectID,
		ListID:    status.ListID,
	}

	if err := s.repository.Create(ctx, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	return s.repository.Update(ctx, payload)
}
