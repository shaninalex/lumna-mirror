package models

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/services/observer"
	"gorm.io/gorm"
)

const (
	EventTaskCreated   observer.Event = "TASK_CREATED"
	EventTaskUpdated   observer.Event = "TASK_UPDATED"
	EventTaskCompleted observer.Event = "TASK_COMPLETED"
)

type Task struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"not null" json:"title"`
	Body      string     `json:"body"`
	Completed bool       `json:"completed"`
	Meta      string     `json:"meta"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *Task) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Task) String() string {
	return fmt.Sprintf("Task id=%d title=%s", s.ID, s.Title)
}
