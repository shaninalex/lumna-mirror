package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Order     uint      `json:"order"`
	Done      bool      `json:"done"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ColumnID  uint      `gorm:"not null;index" json:"column_id"`
	ProjectID uint      `gorm:"not null;index" json:"project_id"`
	BoardID   uint      `gorm:"not null;index" json:"board_id"`
}

func (s *Task) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	s.Done = false
	return nil
}

func (s *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Task) String() string {
	return fmt.Sprintf("Task id=%d title=%s", s.ID, s.Title)
}
