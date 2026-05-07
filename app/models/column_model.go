package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Column struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Order     uint      `json:"order"`
	BoardID   uint      `gorm:"not null;index" json:"board_id"`
	ProjectID uint      `gorm:"not null;index" json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Column) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Column) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Column) String() string {
	return fmt.Sprintf("Column id=%d title=%s", s.ID, s.Title)
}
