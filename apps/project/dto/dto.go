// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// ProjectDto - project dto.
type ProjectDto struct {
	ID        uuid.UUID        `json:"id"`
	Title     string           `json:"title"`
	Code      string           `json:"project_key"`
	Statuses  []*TaskStatusDto `json:"statuses"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// TaskDto - task dto.
type TaskDto struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"creator_id"`
	EpicID      *uuid.UUID `json:"epic_id"`
	SprintID    *uuid.UUID `json:"sprint_id"`
	ProjectID   uuid.UUID  `json:"project_id"`
	Assignee    string     `json:"assignee"`
	Completed   bool       `json:"completed"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StatusID    uuid.UUID  `json:"status"`
	ListIdx     uint       `json:"list_idx"`
	Code        string     `json:"code"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// TaskStatusDto - task status dto.
type TaskStatusDto struct {
	ID          uuid.UUID                `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Complete    bool                     `json:"complete"`
	Index       uint                     `json:"index"`
	Config      *models.TaskStatusConfig `json:"config"`
	//Tasks       []*TaskDto               `json:"tasks"`
}

// ChangeTaskStatusDTO - change task status dto.
type ChangeTaskStatusDTO struct {
	FromStatusID uuid.UUID `json:"from_status"`
	ToStatusID   uuid.UUID `json:"to_status"`
	FromIdx      uint      `json:"from_idx"`
	ToIdx        uint      `json:"to_idx"`
}

type CreateTaskDto struct {
	Title       string    `json:"title"`
	StatusID    uuid.UUID `json:"status_id"`
	ProjectCode string    `json:"project_code"`
}
