package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskEventType string

const (
	TaskEventTypeCreated   TaskEventType = "CREATED"
	TaskEventTypeUpdated   TaskEventType = "UPDATED"
	TaskEventTypeCompleted TaskEventType = "COMPLETED"
)

type TaskEvent struct {
	ID         uint `gorm:"primaryKey"`
	IdentityId uint `gorm:"not null"`
	EventType  TaskEventType
	ValueFrom  string
	ValueTo    string
	CreatedAt  time.Time
}

func (s *TaskEvent) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}
