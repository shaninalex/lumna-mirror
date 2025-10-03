// Copyright © 2025 Lumna. All rights reserved.

package services

import (
	"context"

	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
)

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, projectID uint) ([]*models.Task, error)
	TaskDetail(ctx context.Context, taskID uint) (*models.Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, data *models.Task) error
	TaskUpdate(ctx context.Context, data *models.Task) error
	TaskCreate(ctx context.Context, data *models.Task) (*models.Task, error)
}

type TaskManager interface {
	TaskReader
	TaskWriter
}

type TaskService struct{}

func (t TaskService) TasksList(ctx context.Context, projectID uint) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskDetail(ctx context.Context, taskID uint) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) PatchTaskStatus(ctx context.Context, data *models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskUpdate(ctx context.Context, data *models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskCreate(ctx context.Context, data *models.Task) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskService() *TaskService {
	return &TaskService{}
}
