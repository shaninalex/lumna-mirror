package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID
	User   User

	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Implement IObject interface

func (s *Organization) GetID() uuid.UUID   { return s.ID }
func (s *Organization) SetID(id uuid.UUID) { s.ID = id }

type OrganizationRepository struct {
	Repository[*Organization]
}

func NewOrganizationRepository() *OrganizationRepository {
	s := &OrganizationRepository{}
	return s
}
