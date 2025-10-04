// Copyright © 2025 Lumna. All rights reserved.

package adapter

import (
	"time"

	"gitlab.com/shaninalex/flowreon/domain"
)

// ChangeTaskStatusInput - change task status dto.
type ChangeTaskStatusInput struct {
	FromStatusID uint `json:"from_status"`
	ToStatusID   uint `json:"to_status"`
	FromIdx      uint `json:"from_idx"`
	ToIdx        uint `json:"to_idx"`
}

type TaskDetailInput struct {
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	ListIndex   uint   `json:"list_index"`
	StatusID    uint   `json:"status_id"`
}

type TaskDto struct {
	ID          uint        `json:"id"`
	UserID      uint        `json:"user_id"`
	ProjectID   uint        `json:"project_id"`
	StatusID    uint        `json:"status_id"`
	Title       string      `json:"title"`
	Completed   bool        `json:"completed"`
	Description string      `json:"description"`
	ListIndex   uint        `json:"list_index"`
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
