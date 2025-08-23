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
	Issues         []Issue `gorm:"foreignKey:ProjectID"`

	Title      string
	ProjectKey string `gorm:"uniqueIndex:project_key_uniq"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (s *Project) GetID() uuid.UUID      { return s.ID }
func (s *Project) SetID(id uuid.UUID)    { s.ID = id }
func (s *Project) GetOwner() IUser       { return s.User }
func (s *Project) GetOwnerID() uuid.UUID { return s.UserID }
