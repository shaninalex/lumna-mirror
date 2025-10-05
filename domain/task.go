// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"log"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/db"
	"gitlab.com/shaninalex/flowreon/internal/utils"
)

type Task struct {
	ID          uint
	UserID      uint
	ProjectID   uint
	StatusID    uint
	Title       string
	Completed   bool
	Description *string
	ListIndex   float64
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
	TaskUpdate(ctx context.Context, data *Task) error
	TaskCreate(ctx context.Context, data *Task) (*Task, error)
	TaskDelete(ctx context.Context, taskID uint) error
}

type TaskManager interface {
	TaskReader
	TaskWriter
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

type TaskService struct{}

func (t TaskService) TaskDelete(ctx context.Context, taskID uint) error {
	return db.TaskDelete(ctx, db.GetDb(ctx), taskID)
}

func (t TaskService) TasksList(ctx context.Context, projectID uint) ([]*Task, error) {
	dbTasks, err := db.TaskList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	tasks := make([]*Task, len(dbTasks))
	for i, task := range dbTasks {
		tasks[i] = &Task{
			ID:          task.ID,
			UserID:      task.UserID,
			ProjectID:   task.ProjectID,
			StatusID:    task.StatusID,
			Title:       task.Title,
			Completed:   task.Completed,
			Description: task.Description,
			ListIndex:   task.ListIndex,
			Code:        task.Code,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
		badges, err := db.BadgeTaskList(ctx, db.GetDb(ctx), task.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, badge := range badges {
			tasks[i].Badges = append(tasks[i].Badges, &Badge{
				ID:        badge.ID,
				ProjectID: badge.ProjectID,
				Title:     badge.Title,
				Config:    ToBadgeConfig(badge.Config),
				CreatedAt: badge.CreatedAt,
			})
		}
	}
	return tasks, nil
}

func (t TaskService) TaskDetail(ctx context.Context, taskID uint) (*Task, error) {
	task, err := db.TaskGet(ctx, db.GetDb(ctx), taskID)
	if err != nil {
		return nil, err
	}
	model := &Task{
		ID:          task.ID,
		UserID:      task.UserID,
		ProjectID:   task.ProjectID,
		StatusID:    task.StatusID,
		Title:       task.Title,
		Completed:   task.Completed,
		Description: task.Description,
		ListIndex:   task.ListIndex,
		Code:        task.Code,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	badges, err := db.BadgeTaskList(ctx, db.GetDb(ctx), task.ID)
	if err != nil {
		return nil, err
	}
	for _, badge := range badges {
		model.Badges = append(model.Badges, &Badge{
			ID:        badge.ID,
			ProjectID: badge.ProjectID,
			Title:     badge.Title,
			Config:    ToBadgeConfig(badge.Config),
			CreatedAt: badge.CreatedAt,
		})
	}
	return model, nil
}

func (t TaskService) TaskUpdate(ctx context.Context, data *Task) error {
	return db.TaskUpdate(ctx, db.GetDb(ctx), data.ID, &db.Task{
		Title:       data.Title,
		StatusID:    data.StatusID,
		UserID:      data.UserID,
		Description: data.Description,
		Completed:   data.Completed,
		ListIndex:   data.ListIndex,
	})
}

func (t TaskService) TaskCreate(ctx context.Context, data *Task) (*Task, error) {
	maxIndex := db.TaskGetIndex(ctx, db.GetDb(ctx), data.StatusID)
	task := db.Task{
		UserID:      data.UserID,
		ProjectID:   data.ProjectID,
		StatusID:    data.StatusID,
		Title:       data.Title,
		Code:        utils.GenerateEntityCode("task"),
		Completed:   data.Completed,
		Description: data.Description,
		ListIndex:   maxIndex,
	}
	err := db.TaskSave(ctx, db.GetDb(ctx), &task)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	data.ID = task.ID
	data.CreatedAt = now
	data.UpdatedAt = now
	return data, nil
}
