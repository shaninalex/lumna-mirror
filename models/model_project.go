// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID // Owner
	User   *User

	OrganizationID uuid.UUID `gorm:"uniqueIndex:project_key_uniq"`
	Organization   Organization
	Tasks          []*Task `gorm:"foreignKey:ProjectID"`

	Title      string
	ProjectKey string `gorm:"uniqueIndex:project_key_uniq"`

	Statuses []*TaskStatus `gorm:"foreignKey:ProjectID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetID - returns the id.
func (s *Project) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Project) SetID(id uuid.UUID) { s.ID = id }

// GetOwner - returns the owner.
func (s *Project) GetOwner() AuthUser { return s.User }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetDeletedAt - returns the deleted at.
func (s *Project) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *Project) IsDeleted() bool { return s.DeletedAt.Valid }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
