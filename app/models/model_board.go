package models

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type Board struct {
	ID    EntityId `gorm:"primaryKey" json:"id"`
	Title string   `gorm:"not null" json:"title"`

	ProjectID uint `gorm:"not null;index" json:"project_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *Board) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Board) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Board) String() string {
	return fmt.Sprintf("List id=%d title=%s", s.ID, s.Title)
}
