package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type List struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`

	ProjectID   uint `gorm:"not null;index" json:"project_id"`
	WorkspaceID uint `gorm:"not null;index" json:"workspace_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *List) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *List) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *List) String() string {
	return fmt.Sprintf("List id=%d title=%s", s.ID, s.Title)
}
