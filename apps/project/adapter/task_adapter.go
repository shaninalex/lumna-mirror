// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
)

type TaskCreateInput struct {
	Title string `json:"title"`
}

type TaskDto struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProjectID uint      `json:"project_id"`
	StatusID  uint      `json:"status_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	ListIndex uint      `json:"list_index"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToTaskDto(task *models.Task) *TaskDto {
	return &TaskDto{
		ID:        task.ID,
		UserID:    task.UserID,
		ProjectID: task.ProjectID,
		StatusID:  task.StatusID,
		Title:     task.Title,
		Completed: task.Completed,
		ListIndex: task.ListIndex,
		Code:      task.Code,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func ToTaskListDto(tasks []*models.Task) []*TaskDto {
	output := make([]*TaskDto, len(tasks))
	for i, task := range tasks {
		output[i] = ToTaskDto(task)
	}
	return output
}

type TaskDetailDto struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	ProjectID   uint      `json:"project_id"`
	StatusID    uint      `json:"status_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	ListIndex   uint      `json:"list_index"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
