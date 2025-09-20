// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
)

// Organization - organization.
type Organization struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *Organization) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *Organization) SetID(id uuid.UUID) { s.ID = id }

// GetOwnerID - returns the owner id.
func (s *Organization) GetOwnerID() uuid.UUID { return s.UserID }

// IsOwner - checks if it is owner.
func (s *Organization) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Organization) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Organization) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Organization) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
