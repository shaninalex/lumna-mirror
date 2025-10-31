// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"github.com/shaninalex/lumna/app/domain"
)

type TaskCreateInput struct {
	Title string `json:"title"`
}

type TaskDto struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"user_id"`
	ProjectID int64       `json:"project_id"`
	StatusID  int64       `json:"status_id"`
	Title     string      `json:"title"`
	Completed bool        `json:"completed"`
	ListIndex float64     `json:"list_index"`
	Code      string      `json:"code"`
	Badges    []*BadgeDto `json:"badges"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func ToTaskDto(task *domain.Task) *TaskDto {
	badges := make([]*BadgeDto, len(task.Badges))
	for i, badge := range task.Badges {
		badges[i] = NewBadgeDto(badge)
	}
	return &TaskDto{
		ID:        task.ID,
		UserID:    task.UserID,
		ProjectID: task.ProjectID,
		StatusID:  task.StatusID,
		Title:     task.Title,
		Completed: task.Completed,
		ListIndex: task.ListIndex,
		Code:      task.Code,
		Badges:    badges,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func ToTaskListDto(tasks []*domain.Task) []*TaskDto {
	output := make([]*TaskDto, len(tasks))
	for i, task := range tasks {
		output[i] = ToTaskDto(task)
	}
	return output
}
