package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`
	Order uint   `json:"order"`

	ListID     uint `gorm:"not null;index" json:"board_id"`
	ProjectID   uint `gorm:"not null;index" json:"project_id"`
	WorkspaceID uint `gorm:"not null;index" json:"workspace_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Status) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Status) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Status) String() string {
	return fmt.Sprintf("Column id=%d title=%s", s.ID, s.Title)
}
