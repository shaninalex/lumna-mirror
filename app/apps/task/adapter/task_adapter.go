// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/domain"
)

// ChangeTaskStatusInput - change task status dto.
type ChangeTaskStatusInput struct {
	FromStatusID int64   `json:"from_status"`
	ToStatusID   int64   `json:"to_status"`
	FromIdx      int64   `json:"from_idx"`
	ToIdx        float64 `json:"to_idx"`
}

type TaskDetailInput struct {
	Title       string  `json:"title"`
	Completed   bool    `json:"completed"`
	Description string  `json:"description"`
	ListIndex   float64 `json:"list_index"`
	StatusID    int64   `json:"status_id"`
}

type TaskDto struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"user_id"`
	ProjectID   int64       `json:"project_id"`
	StatusID    int64       `json:"status_id"`
	Title       string      `json:"title"`
	Completed   bool        `json:"completed"`
	Description string      `json:"description"`
	ListIndex   float64     `json:"list_index"`
	Code        string      `json:"code"`
	Badges      []*BadgeDto `json:"badges"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func ToTaskDto(task *domain.Task) *TaskDto {
	badges := make([]*BadgeDto, len(task.Badges))
	for i, badge := range task.Badges {
		badges[i] = NewBadgeDto(badge)
	}
	out := &TaskDto{
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

	if task.Description != nil {
		out.Description = *task.Description
	}

	return out
}
