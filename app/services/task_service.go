package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type TaskService interface {
	List(ctx context.Context, query ServiceTaskListQuery) ([]models.Task, []models.BoardTask, error)
	GetTask(ctx context.Context, taskID uint) (*models.Task, error)
	ReorderTask(ctx context.Context, taskID uint, listListID uint, order uint) error
	CreateTask(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error)
	UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error
}

type taskService struct {
	repository repositories.TaskRepository
	bus        observer.Observer
}

var _ TaskService = (*taskService)(nil)

func NewTaskService(
	repository repositories.TaskRepository,
	bus observer.Observer,
) TaskService {
	return &taskService{
		repository: repository,
		bus:        bus,
	}
}

type ServiceTaskListQuery struct {
	BoardId uint `form:"board_id"`
}

func (s *taskService) List(ctx context.Context, query ServiceTaskListQuery) ([]models.Task, []models.BoardTask, error) {
	bts, err := s.repository.GetTaskBoards(ctx, query.BoardId)
	if err != nil {
		return nil, nil, err
	}

	tasks, err := s.repository.GetTasksByBoardId(ctx, query.BoardId)
	if err != nil {
		return nil, nil, err
	}

	return tasks, bts, nil

}

func (s *taskService) GetTask(ctx context.Context, taskID uint) (*models.Task, error) {
	return s.repository.GetByID(ctx, taskID)
}

func (s *taskService) ReorderTask(ctx context.Context, taskID uint, listListID uint, order uint) error {
	return s.repository.Reorder(ctx, taskID, listListID, order)
}

func (s *taskService) CreateTask(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error) {
	task, err := s.repository.AddTaskToBoard(ctx, payload)
	if err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventTaskCreated, task)
	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID uint, payload *models.Task) error {
	if _, err := s.GetTask(ctx, taskID); err != nil {
		return err
	}
	return s.repository.Update(ctx, payload)
}
