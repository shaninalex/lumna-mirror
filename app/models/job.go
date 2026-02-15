package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID          uint `gorm:"primaryKey"`
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
	return nil
}
