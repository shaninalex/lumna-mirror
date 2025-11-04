package db

import (
	"time"
)

// Comment - task comment:w
type Comment struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	TaskID    int64     `db:"project_id"`
	Content   string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

// GetID - returns the id.
func (s *Comment) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *Comment) SetID(id int64) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Comment) GetOwnerID() int64 { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Comment) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Comment) GetCreatedAt() time.Time { return s.CreatedAt }

// GetCreatedBy - returns the created by.
func (s *Comment) GetCreatedBy() int64 { return s.UserID }
