package database

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
