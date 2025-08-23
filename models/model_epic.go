package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (s *Epic) GetID() uuid.UUID             { return s.ID }
func (s *Epic) SetID(id uuid.UUID)           { s.ID = id }
func (s *Epic) GetOwner() AuthUser           { return s.User }
func (s *Epic) GetOwnerID() uuid.UUID        { return s.UserID }
func (s *Epic) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }
func (s *Epic) GetCreatedAt() time.Time      { return s.CreatedAt }
func (s *Epic) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}
func (s *Epic) IsDeleted() bool         { return s.DeletedAt.Valid }
func (s *Epic) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
