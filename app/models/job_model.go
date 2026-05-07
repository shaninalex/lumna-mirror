package models

import (
	"time"

	"gorm.io/gorm"
)

type JobStatus string

var (
	// JobStatusPending - waiting for execution
	JobStatusPending = JobStatus("pending")

	// JobStatusRunning - currently executing
	JobStatusRunning = JobStatus("running")

	// JobStatusSuccess - completed without errors
	JobStatusSuccess = JobStatus("success")

	// JobStatusRepeat after 3-rd attempt job becomes with error start
	JobStatusRepeat = JobStatus("repeat")

	// JobStatusError dead end, executed with error
	JobStatusError = JobStatus("error")
)

type Job struct {
	ID          uint `gorm:"primaryKey"`
	Type        string
	Payload     string
	Status      JobStatus
	Attempts    uint
	AvailableAt time.Time // when to start? Delete this for now?
	LockedAt    *time.Time
	LockedBy    *string // reason why job is locked
	CreatedAt   time.Time
}

func (s *Job) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}
