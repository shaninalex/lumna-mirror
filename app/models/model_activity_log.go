package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ActivityLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Summary    string    `gorm:"not null" json:"summary"`
	IdentityID uint      `gorm:"not null;index" json:"identity_id"`
	Identity   Identity  `gorm:"foreignKey:IdentityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EntityID   uint      `gorm:"not null;index" json:"entity_id"`
	EntityType string    `gorm:"not null;index" json:"entity_type"`
	Action     string    `gorm:"type:text;not null;index" json:"action"`
	CreatedAt  time.Time `json:"created_at"`
}

func (s *ActivityLog) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	// build summary
	return nil
}

func (s *ActivityLog) String() string {
	return fmt.Sprintf("Project id=%d title=%s", s.ID, s.Summary)
}
