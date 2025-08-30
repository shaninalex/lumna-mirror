// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (s *Organization) GetID() uuid.UUID             { return s.ID }
func (s *Organization) SetID(id uuid.UUID)           { s.ID = id }
func (s *Organization) GetOwner() AuthUser           { return s.User }
func (s *Organization) GetOwnerID() uuid.UUID        { return s.UserID }
func (s *Organization) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }
func (s *Organization) GetCreatedAt() time.Time      { return s.CreatedAt }
func (s *Organization) GetUpdatedAt() time.Time      { return s.UpdatedAt }
func (s *Organization) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}
func (s *Organization) IsDeleted() bool         { return s.DeletedAt.Valid }
func (s *Organization) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
