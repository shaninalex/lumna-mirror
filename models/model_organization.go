// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Organization - organization.
type Organization struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	// Creator of an Organization
	UserID uuid.UUID
	User   *User

	Title       string
	Description string

	Users []*User `gorm:"foreignKey:OrganizationID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetID - returns the id.
func (s *Organization) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Organization) SetID(id uuid.UUID) { s.ID = id }

// GetOwner - returns the owner.
func (s *Organization) GetOwner() AuthUser { return s.User }

// GetOwnerID - returns the owner id.
func (s *Organization) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Organization) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Organization) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Organization) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetDeletedAt - returns the deleted at.
func (s *Organization) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *Organization) IsDeleted() bool { return s.DeletedAt.Valid }

// GetCreatedBy - returns the created by.
func (s *Organization) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
