package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title   string    `gorm:"not null"`
	Content string

	IdentityID uuid.UUID `gorm:"type:uuid"`
	Shared     bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
