// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
)

// Project - project.
type Project struct {
	ID             uuid.UUID `db:"id"`
	UserID         uuid.UUID `db:"user_id"`
	OrganizationID uuid.UUID `db:"organization_id"`
	Title          string    `db:"title"`
	Code           string    `db:"code"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *Project) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Project) SetID(id uuid.UUID) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
