package database

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

// Implement IObject interface

func (s *Sprint) GetID() uuid.UUID   { return s.ID }
func (s *Sprint) SetID(id uuid.UUID) { s.ID = id }

type SprintRepository struct {
	Repository[*Sprint]
}

func NewSprintRepository() *SprintRepository {
	s := &SprintRepository{}
	return s
}
