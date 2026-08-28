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
	ID        EntityId `gorm:"primaryKey"`
	Title     string   `gorm:"not null"`
	Body      string
	Completed bool
	Meta      string
	ProjectId uint       `gorm:"-"`
	BoardIds  []EntityId `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt *time.Time
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

type TaskCreateOnBoard struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Position *uint  `json:"position"`
	BoardId  *uint  `json:"board_id"`
	ColumnId *uint  `json:"column_id"`
}

type TaskCreateBacklog struct {
	Title string `json:"title"`
}
