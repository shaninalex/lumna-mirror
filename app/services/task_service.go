package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

type TaskService interface {
	List(ctx context.Context, query ServiceTaskListQuery) ([]models.Task, error)
	GetTask(ctx context.Context, taskID int) (*models.Task, error)
	CreateTask(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error)
	UpdateTask(ctx context.Context, taskID int, payload *models.Task) error

	Move(ctx context.Context, boardId int, rearange models.RearangeTask) error
	Transfer(ctx context.Context, transfer models.TransferTaskBetweenColumns) error
}

type taskService struct {
	repository        repositories.TaskRepository
	storageRepository repositories.TaskStorageRepository
	bus               observer.Observer
}

var _ TaskService = (*taskService)(nil)

func NewTaskService(
	repository repositories.TaskRepository,
	storageRepository repositories.TaskStorageRepository,
	bus observer.Observer,
) TaskService {
	return &taskService{
		repository:        repository,
		storageRepository: storageRepository,
		bus:               bus,
	}
}

type ServiceTaskListQuery struct {
	BoardId int `form:"board_id"`
}

func (s *taskService) List(ctx context.Context, query ServiceTaskListQuery) ([]models.Task, error) {
	return s.storageRepository.Filter(ctx, query.BoardId)
}

func (s *taskService) GetTask(ctx context.Context, taskID int) (*models.Task, error) {
	return s.repository.GetByID(ctx, taskID)
}

func (s *taskService) CreateTask(ctx context.Context, payload *models.TaskCreateOnBoard) (*models.Task, error) {
	task := &models.Task{
		Title:     payload.Title,
		Body:      payload.Body,
		ProjectId: payload.ProjectId,
		Boards: []models.TaskBoard{
			{
				Position: payload.Position,
				BoardId:  payload.BoardId,
				ColumnId: payload.ColumnId,
			},
		},
	}
	err := s.storageRepository.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, models.EventTaskCreated, task)
	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID int, payload *models.Task) error {
	if _, err := s.GetTask(ctx, taskID); err != nil {
		return err
	}
	return s.repository.Update(ctx, payload)
}

// Move implements [TaskService].
func (s *taskService) Move(ctx context.Context, boardId int, rearange models.RearangeTask) error {
	for i, t := range rearange.Tasks {
		if err := s.storageRepository.MovePosition(ctx, t, boardId, rearange.ColumnId, i); err != nil {
			return err
		}
	}

	return nil
}

// Transfer implements [TaskService].
func (s *taskService) Transfer(ctx context.Context, transfer models.TransferTaskBetweenColumns) error {
	for i, t := range transfer.From.Tasks {
		if err := s.storageRepository.MovePosition(ctx, t, transfer.BoardId, transfer.From.ColumnId, i); err != nil {
			return err
		}
	}

	for i, t := range transfer.To.Tasks {
		if err := s.storageRepository.MovePosition(ctx, t, transfer.BoardId, transfer.To.ColumnId, i); err != nil {
			return err
		}
	}

	return nil
}
