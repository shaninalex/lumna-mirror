// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
)

// Task - task.
type Task struct {
	ID             uuid.UUID  `db:"id"`
	UserID         uuid.UUID  `db:"user_id"`
	EpicID         *uuid.UUID `db:"epic_id"`
	OrganizationID uuid.UUID  `db:"organization_id"`
	SprintID       *uuid.UUID `db:"sprint_id"`
	ProjectID      uuid.UUID  `db:"project_id"`
	TaskStatusID   uuid.UUID  `db:"tas_status_id"`
	Assignee       string     `db:"assignee"`
	Title          string     `db:"title"`
	Completed      bool       `db:"completed"`
	Description    string     `db:"description"`
	ListIndex      uint       `db:"list_index"`
	Code           string     `db:"code"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
}

// GetID - returns the id.
func (s *Task) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Task) SetID(id uuid.UUID) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Task) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Task) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Task) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Task) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Task) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }

// SetCode - sets the code.
func (s *Task) SetCode(code string) { s.Code = code }

// GetCode - returns the code.
func (s *Task) GetCode() string { return s.Code }
