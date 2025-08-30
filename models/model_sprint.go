// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sprint struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID
	User   *User

	OrganizationID uuid.UUID
	Organization   *Organization

	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetID - returns the id.
func (s *Sprint) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Sprint) SetID(id uuid.UUID) { s.ID = id }

// GetOwner - returns the owner.
func (s *Sprint) GetOwner() AuthUser { return s.User }

// GetOwnerID - returns the owner id.
func (s *Sprint) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Sprint) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Sprint) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Sprint) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetDeletedAt - returns the deleted at.
func (s *Sprint) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *Sprint) IsDeleted() bool { return s.DeletedAt.Valid }

// GetCreatedBy - returns the created by.
func (s *Sprint) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
