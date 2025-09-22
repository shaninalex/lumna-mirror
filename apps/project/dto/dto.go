// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"time"

	"gitlab.com/shaninalex/flowreon/models"
)

// ProjectDto - project dto.
type ProjectDto struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Code      string           `json:"project_key"`
	Statuses  []*TaskStatusDto `json:"statuses"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// TaskDto - task dto.
type TaskDto struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"creator_id"`
	ProjectID   uint       `json:"project_id"`
	Assignee    string     `json:"assignee"`
	Completed   bool       `json:"completed"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StatusID    uint       `json:"status"`
	ListIdx     uint       `json:"list_idx"`
	Code        string     `json:"code"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// TaskStatusDto - task status dto.
type TaskStatusDto struct {
	ID          uint                     `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Complete    bool                     `json:"complete"`
	Index       uint                     `json:"index"`
	Config      *models.TaskStatusConfig `json:"config"`
	//Tasks       []*TaskDto               `json:"tasks"`
}

// ChangeTaskStatusDTO - change task status dto.
type ChangeTaskStatusDTO struct {
	FromStatusID uint `json:"from_status"`
	ToStatusID   uint `json:"to_status"`
	FromIdx      uint `json:"from_idx"`
	ToIdx        uint `json:"to_idx"`
}

type CreateTaskDto struct {
	Title       string `json:"title"`
	StatusID    uint   `json:"status_id"`
	ProjectCode string `json:"project_code"`
}
