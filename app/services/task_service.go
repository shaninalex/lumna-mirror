package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetTask(ctx context.Context, taskID uuid.UUID) (*models.Task, error) {
	database := db.GetDB(ctx)
	task := &models.Task{}
	if result := database.Where("id = ?", taskID).First(&task); result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (s *TaskService) ReorderTask(ctx context.Context, taskID uuid.UUID, boardListID uuid.UUID, order uint) error {
	database := db.GetDB(ctx)
	return database.Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]any{
			"board_list_id": boardListID,
			"order":         order,
		}).Error
}
