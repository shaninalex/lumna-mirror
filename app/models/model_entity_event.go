package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityEventType string

const (
	EntityEventTaskCreated   EntityEventType = "task/created"
	EntityEventTaskUpdated   EntityEventType = "task/updated"
	EntityEventTaskCompleted EntityEventType = "task/completed"
)

type EntityEvent struct {
	ID         int  `gorm:"primaryKey"`
	IdentityId *int `gorm:"null"`
	EntityId   *int
	EntityType *string
	EventType  EntityEventType
	Data       string
	CreatedAt  time.Time
}

func (s EntityEvent) TableName() string {
	return "entity_events"
}

func (s *EntityEvent) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}
