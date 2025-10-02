// Copyright © 2025 Lumna. All rights reserved.

package models

import (
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
