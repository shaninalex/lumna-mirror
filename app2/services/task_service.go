package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
	"gitlab.com/shaninalex/lumna/app2/repositories"
)

type TaskManager interface {
	Get(ctx context.Context, id uint) (*models.Task, error)
	BoardTasks(ctx context.Context, boardId uint) ([]*models.Task, error)
	Create(ctx context.Context, entry *models.Task) error
	Delete(ctx context.Context, id uint) error
	Patch(ctx context.Context, id uint, opts db.SetExpr) error
}

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskManager() TaskManager {
	return &TaskService{
		taskRepository: repositories.NewTaskRepository(),
	}
}

var _ ProjectManager = (*ProjectService)(nil)

func (s *TaskService) Get(ctx context.Context, id uint) (*models.Task, error) {
	return s.taskRepository.Get(ctx, id)
}

func (s *TaskService) BoardTasks(ctx context.Context, boardId uint) ([]*models.Task, error) {
	return s.taskRepository.List(ctx, db.Eq("board_id", boardId))
}

func (s *TaskService) Create(ctx context.Context, entry *models.Task) error {
	return s.taskRepository.Create(ctx, entry)
}

func (s *TaskService) Delete(ctx context.Context, taskId uint) error {
	return s.taskRepository.Delete(ctx, taskId)
}

func (s *TaskService) Patch(ctx context.Context, id uint, opts db.SetExpr) error {
	return s.taskRepository.Update(ctx, id, opts)
}

type TaskPayloadModel struct {
	Name   string `json:"name"`
	ListId uint   `json:"list_id"`
	Order  uint   `json:"order"`
}
