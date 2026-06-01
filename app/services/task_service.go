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

type ServiceTaskListQuery struct {
	ProjectID *uint `form:"project_id,omitempty"`
}

func (s *TaskService) List(ctx context.Context, query ServiceTaskListQuery) ([]*models.Task, error) {
	return s.repository.List(ctx, repositories.TaskListQuery{
		ProjectID: query.ProjectID,
	})
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
	StatusID  uint   `json:"status_id"`
}

func (s *TaskService) CreateTask(ctx context.Context, payload *TaskPayload) (*models.Task, error) {
	status, err := s.statusRepository.GetByID(ctx, payload.StatusID)
	if err != nil {
		return nil, err
	}
	task := models.Task{
		Title:     payload.Title,
		Order:     payload.Order,
		StatusID:  status.ID,
		ProjectID: status.ProjectID,
	}

	if err = s.repository.Create(ctx, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	return s.repository.Update(ctx, payload)
}
