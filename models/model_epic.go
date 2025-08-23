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

func (s *Epic) GetID() uuid.UUID      { return s.ID }
func (s *Epic) SetID(id uuid.UUID)    { s.ID = id }
func (s *Epic) GetOwner() IUser       { return s.User }
func (s *Epic) GetOwnerID() uuid.UUID { return s.UserID }
