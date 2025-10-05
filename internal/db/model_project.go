// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"time"
)

// Project - project.
type Project struct {
	ID        uint      `db:"id"`
	UserID    uint      `db:"user_id"`
	Title     string    `db:"title"`
	Code      string    `db:"code"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *Project) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *Project) SetID(id uint) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() uint { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() uint { return s.GetOwnerID() }
