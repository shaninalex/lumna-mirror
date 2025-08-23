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

func (s *Organization) GetID() uuid.UUID      { return s.ID }
func (s *Organization) SetID(id uuid.UUID)    { s.ID = id }
func (s *Organization) GetOwner() IUser       { return s.User }
func (s *Organization) GetOwnerID() uuid.UUID { return s.UserID }
