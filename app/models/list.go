package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardList struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string    `gorm:"not null" json:"title"`
	Order uint      `json:"order"`

	BoardID uuid.UUID `gorm:"type:uuid;not null;index" json:"board_id"`

	Tasks []Task

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *BoardList) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *BoardList) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *BoardList) String() string {
	return fmt.Sprintf("BoardList id=%s title=%s", s.ID.String(), s.Title)
}
