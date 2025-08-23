package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Epic struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID
	User   User

	ProjectID uuid.UUID
	Project   Project

	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Implement IObject interface

func (s *Epic) GetID() uuid.UUID   { return s.ID }
func (s *Epic) SetID(id uuid.UUID) { s.ID = id }

type EpicRepository struct {
	Repository[*Epic]
}

func NewEpicRepository() *EpicRepository {
	s := &EpicRepository{}
	return s
}
