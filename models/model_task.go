// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID
	User   *User

	EpicID *uuid.UUID
	Epic   *Epic

	OrganizationID uuid.UUID
	Organization   *Organization

	SprintID *uuid.UUID
	Sprint   *Sprint

	ProjectID uuid.UUID
	Project   *Project

	TaskStatusID uuid.UUID
	TaskStatus   *TaskStatus `gorm:"foreignKey:TaskStatusID;references:ID"`

	// Assignee who is assigned to do that task. NOTE: Should be issue_assignee separate relation.
	Assignee string

	Title       string
	Completed   bool
	Description string

	// ListIndex - where in a list of statuses it's currently in
	ListIndex uint

	// Code - short task code like "task-123123".
	Code string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetID - returns the id.
func (s *Task) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Task) SetID(id uuid.UUID) { s.ID = id }

// GetOwner - returns the owner.
func (s *Task) GetOwner() AuthUser { return s.User }

// GetOwnerID - returns the owner id.
func (s *Task) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Task) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Task) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Task) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetDeletedAt - returns the deleted at.
func (s *Task) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *Task) IsDeleted() bool { return s.DeletedAt.Valid }

// GetCreatedBy - returns the created by.
func (s *Task) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }

// SetCode - sets the code.
func (s *Task) SetCode(code string) { s.Code = code }

// GetCode - returns the code.
func (s *Task) GetCode() string { return s.Code }
