package services

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
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
	repository        repositories.TaskRepository
	statusRepository  repositories.StatusRepository
	projectRepository repositories.ProjectRepository
}

var _ TaskService = (*taskService)(nil)

func NewTaskService(
	repository repositories.TaskRepository,
	statusRepository repositories.StatusRepository,
	projectRepository repositories.ProjectRepository,
) TaskService {
	return &taskService{
		repository:        repository,
		statusRepository:  statusRepository,
		projectRepository: projectRepository,
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

	if err := s.taskCode(ctx, &task, payload.ProjectID); err != nil {
		return nil, err
	}

	if err := s.repository.Create(ctx, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	return s.repository.Update(ctx, payload)
}

var (
	ErrorTaskUnableGenerateCode = errors.New("unable to generate task code")
	GenerateTaskCodeMaxAttempts = 25
)

// taskCode - sequential task code generator
// NOTE: not sure about that logic
func (s *taskService) taskCode(ctx context.Context, task *models.Task, projectId uint) error {
	project, err := s.projectRepository.GetByID(ctx, projectId)
	if err != nil {
		return err
	}
	tasks, err := s.repository.List(ctx, repositories.TaskListQuery{
		ProjectID: &projectId,
	})

	for i := 1; i <= GenerateTaskCodeMaxAttempts; i++ {
		code := utils.TaskCode(project.Title, len(tasks)+i)
		tasks, err = s.repository.List(ctx, repositories.TaskListQuery{
			ProjectID: &projectId,
			Code:      &code,
		})
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			task.Code = code
			return nil
		}
	}

	return ErrorTaskUnableGenerateCode
}
