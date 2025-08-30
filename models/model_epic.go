// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Epic - epic.
type Epic struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID
	User   *User

	ProjectID uuid.UUID
	Project   *Project

	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// GetID - returns the id.
func (s *Epic) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Epic) SetID(id uuid.UUID) { s.ID = id }

// GetOwner - returns the owner.
func (s *Epic) GetOwner() AuthUser { return s.User }

// GetOwnerID - returns the owner id.
func (s *Epic) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Epic) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Epic) GetCreatedAt() time.Time { return s.CreatedAt }

// GetDeletedAt - returns the deleted at.
func (s *Epic) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *Epic) IsDeleted() bool { return s.DeletedAt.Valid }

// GetCreatedBy - returns the created by.
func (s *Epic) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
