package domain

import (
	"time"
)

const EntityTypeTask = "task"

type Task struct {
	Id          int64
	UserID      int64
	ProjectID   int64
	StatusID    int64
	Title       string
	Completed   bool
	Description *string
	ListIndex   float64
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// related structures
	Badges        []*Badge
	Comments      []*Comment
	CommentsCount int
}

// GetID - returns the id.
func (s *Task) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Task) SetID(id int64) { s.Id = id }

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
