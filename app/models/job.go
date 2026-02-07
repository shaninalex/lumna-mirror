package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type        string
	Payload     string
	Status      string
	Attempts    uint
	AvailableAt time.Time
	LockedAt    time.Time
	LockedBy    string // reason why job is locked
	CreatedAt   time.Time
}

func (u *Job) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
