package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string    `gorm:"not null" json:"title"`
	Order uint      `json:"order"`
	Done  bool      `json:"done"`

	ProjectID uuid.UUID `gorm:"type:uuid;not null;index" json:"project_id"`
	ListID    uuid.UUID `gorm:"type:uuid;not null;index" json:"list_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Task) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
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
	return fmt.Sprintf("Task id=%s title=%s", s.ID.String(), s.Title)
}
