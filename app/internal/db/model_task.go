// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"time"
)

// Task - task.
type Task struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	ProjectID   int64     `db:"project_id"`
	StatusID    int64     `db:"status_id"`
	Title       string    `db:"title"`
	Completed   bool      `db:"completed"`
	Description *string   `db:"description"`
	ListIndex   float64   `db:"list_index"`
	Code        string    `db:"code"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *Task) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *Task) SetID(id int64) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Task) GetOwnerID() int64 { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Task) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Task) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Task) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Task) GetCreatedBy() int64 { return s.GetOwnerID() }

// SetCode - sets the code.
func (s *Task) SetCode(code string) { s.Code = code }

// GetCode - returns the code.
func (s *Task) GetCode() string { return s.Code }
