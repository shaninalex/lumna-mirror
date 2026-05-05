package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type TaskService struct {
	repository       repositories.TaskRepository
	columnRepository repositories.ColumnRepository
}

func NewTaskService(
	repository repositories.TaskRepository,
	columnRepository repositories.ColumnRepository,
) *TaskService {
	return &TaskService{
		repository:       repository,
		columnRepository: columnRepository,
	}
}

func (s *TaskService) GetTask(ctx context.Context, taskID uint) (*models.Task, error) {
	return s.repository.GetByID(ctx, taskID)
}

func (s *TaskService) ReorderTask(ctx context.Context, taskID uint, boardListID uint, order uint) error {
	return s.repository.Reorder(ctx, taskID, boardListID, order)
}

// TaskPayload - used to create/partial update task
// TODO: add validators
type TaskPayload struct {
	Title     string `json:"title"`
	Order     uint   `json:"order"`
	ProjectID uint   `json:"project_id"`
	ColumnID  uint   `json:"column_id"`
}

func (s *TaskService) CreateTask(ctx context.Context, payload *TaskPayload) (*models.Task, error) {
	column, err := s.columnRepository.GetByID(ctx, payload.ColumnID)
	if err != nil {
		return nil, err
	}
	task := models.Task{
		Title:     payload.Title,
		Order:     payload.Order,
		ColumnID:  column.ID,
		ProjectID: column.ProjectID,
		BoardID:   column.BoardID,
	}

	if err := s.repository.Create(ctx, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	return s.repository.Update(ctx, payload)
}
