package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type TaskService interface {
	List(ctx context.Context, query ServiceTaskListQuery) ([]*models.Task, error)
	GetTask(ctx context.Context, taskID uint) (*models.Task, error)
	ReorderTask(ctx context.Context, taskID uint, listListID uint, order uint) error
	CreateTask(ctx context.Context, payload *TaskPayload) (*models.Task, error)
	UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error
}

type taskService struct {
	repository       repositories.TaskRepository
	statusRepository repositories.StatusRepository
}

var _ TaskService = (*taskService)(nil)

func NewTaskService(
	repository repositories.TaskRepository,
	statusRepository repositories.StatusRepository,
) TaskService {
	return &taskService{
		repository:       repository,
		statusRepository: statusRepository,
	}
}

type ServiceTaskListQuery struct {
	ProjectID *uint `form:"project_id,omitempty"`
}

func (s *taskService) List(ctx context.Context, query ServiceTaskListQuery) ([]*models.Task, error) {
	return s.repository.List(ctx, repositories.TaskListQuery{
		ProjectID: query.ProjectID,
	})
}

func (s *taskService) GetTask(ctx context.Context, taskID uint) (*models.Task, error) {
	return s.repository.GetByID(ctx, taskID)
}

func (s *taskService) ReorderTask(ctx context.Context, taskID uint, listListID uint, order uint) error {
	return s.repository.Reorder(ctx, taskID, listListID, order)
}

// TaskPayload - used to create/partial update task
// TODO: add validators
type TaskPayload struct {
	Title     string `json:"title"`
	ProjectID uint   `json:"project_id"`
	Order     *uint  `json:"order"`
	StatusID  *uint  `json:"status_id"`
}

func (s *taskService) CreateTask(ctx context.Context, payload *TaskPayload) (*models.Task, error) {
	task := models.Task{
		Title:     payload.Title,
		ProjectID: payload.ProjectID,
	}
	if payload.StatusID != nil {
		status, err := s.statusRepository.GetByID(ctx, *payload.StatusID)
		if err != nil {
			return nil, err
		}
		task.StatusID = &status.ID
	}

	if err := s.repository.Create(ctx, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	return s.repository.Update(ctx, payload)
}
