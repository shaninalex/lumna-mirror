package models

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type Status struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`

	// Order - the order of status (column) in board view of list
	// NOTE: it make sence only if this view. May be create `meta` column
	// instead and use this information there? With all other like color,
	// icon, expanded, collapsed and so on...
	Order uint `json:"order"`

	ProjectID uint `gorm:"not null;index" json:"project_id"`
	ListId    uint `gorm:"not null;index" json:"list_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *Status) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Status) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Status) String() string {
	return fmt.Sprintf("Column id=%d title=%s", s.ID, s.Title)
}

type StatusMeta struct {
	Order    uint   `json:"order"` // NOTE: remove from base model and move here
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Expanded bool   `json:"expanded"`
}
