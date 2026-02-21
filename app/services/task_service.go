package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) GetTask(ctx context.Context, taskID uint) (*models.Task, error) {
	task := &models.Task{}
	if result := s.db.WithContext(ctx).Where("id = ?", taskID).First(&task); result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (s *TaskService) ReorderTask(ctx context.Context, taskID uint, boardListID uint, order uint) error {
	return s.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]any{
			"board_list_id": boardListID,
			"order":         order,
		}).Error
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
	column := &models.Column{}
	result := s.db.WithContext(ctx).Where("id = ?", payload.ColumnID).First(&column)
	if result.Error != nil {
		return nil, result.Error
	}
	task := models.Task{
		Title:     payload.Title,
		Order:     payload.Order,
		ColumnID:  column.ID,
		ProjectID: column.ProjectID,
	}
	if result := s.db.WithContext(ctx).Create(&task); result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}
