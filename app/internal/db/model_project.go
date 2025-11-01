package db

import (
	"time"
)

// Project - project.
type Project struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Code      string    `db:"code"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *Project) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *Project) SetID(id int64) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() int64 { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() int64 { return s.GetOwnerID() }
