// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"time"
)

type Task struct {
	ID          uint
	UserID      uint
	ProjectID   uint
	StatusID    uint
	Title       string
	Completed   bool
	Description *string
	ListIndex   uint
	Code        string
	Badges      []*Badge
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TaskReader - task reader.
type TaskReader interface {
	TasksList(ctx context.Context, projectID uint) ([]*Task, error)
	TaskDetail(ctx context.Context, taskID uint) (*Task, error)
}

// TaskWriter - task writer.
type TaskWriter interface {
	PatchTaskStatus(ctx context.Context, data *Task) error
	TaskUpdate(ctx context.Context, data *Task) error
	TaskCreate(ctx context.Context, data *Task) (*Task, error)
}

type TaskManager interface {
	TaskReader
	TaskWriter
}

type TaskService struct{}

func (t TaskService) TasksList(ctx context.Context, projectID uint) ([]*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskDetail(ctx context.Context, taskID uint) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) PatchTaskStatus(ctx context.Context, data *Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskUpdate(ctx context.Context, data *Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) TaskCreate(ctx context.Context, data *Task) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskService() *TaskService {
	return &TaskService{}
}
