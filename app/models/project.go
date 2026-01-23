package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string    `gorm:"not null" json:"title"`

	Boards []Board `json:"boards"`
	Tasks  []Task  `json:"tasks"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Project) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return nil
}

func (s *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Project) String() string {
	return fmt.Sprintf("Project id=%s title=%s", s.ID.String(), s.Title)
}
