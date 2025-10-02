// Copyright © 2025 Lumna. All rights reserved.

package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	ProjectID   uint      `json:"project_id"`
	StatusID    uint      `json:"status_id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description *string   `json:"description"`
	ListIndex   uint      `json:"list_index"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
