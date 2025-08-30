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

func (s *Task) GetID() uuid.UUID             { return s.ID }
func (s *Task) SetID(id uuid.UUID)           { s.ID = id }
func (s *Task) GetOwner() AuthUser           { return s.User }
func (s *Task) GetOwnerID() uuid.UUID        { return s.UserID }
func (s *Task) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }
func (s *Task) GetCreatedAt() time.Time      { return s.CreatedAt }
func (s *Task) GetUpdatedAt() time.Time      { return s.UpdatedAt }
func (s *Task) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}
func (s *Task) IsDeleted() bool         { return s.DeletedAt.Valid }
func (s *Task) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
func (s *Task) SetCode(code string)     { s.Code = code }
func (s *Task) GetCode() string         { return s.Code }
