package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID // Owner
	User   User

	OrganizationID uuid.UUID
	Organization   Organization

	Title string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Implement IObject interface

func (s *Project) GetID() uuid.UUID   { return s.ID }
func (s *Project) SetID(id uuid.UUID) { s.ID = id }

type ProjectRepository struct {
	Repository[*Project]
}

func NewProjectRepository() *ProjectRepository {
	s := &ProjectRepository{}
	return s
}
