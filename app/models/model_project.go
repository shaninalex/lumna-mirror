package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type Project struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Key         string     `gorm:"not null" json:"key"`
	WorkspaceID uint       `gorm:"not null;index" json:"workspace_id"`
	OwnerID     *uint      `gorm:"null;index" json:"owner_id"`
	Meta        *string    `gorm:"meta" json:"meta"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (s *Project) GetTitle() string {
	return s.Title
}

func (s *Project) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return nil
}

func (s *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Project) String() string {
	return fmt.Sprintf("Project id=%d title=%s", s.ID, s.Title)
}

type ProjectMeta struct {
}

func NewProjectMeta() *ProjectMeta {
	return &ProjectMeta{}
}

func (s *Project) GetMeta() *ProjectMeta {
	m := NewProjectMeta()
	if s.Meta == nil {
		return m
	}
	if err := json.Unmarshal([]byte(*s.Meta), m); err != nil {
		return NewProjectMeta()
	}
	return m
}

func (s *Project) SetMeta(m *ProjectMeta) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	s.Meta = utils.Pointer(string(b))
	return nil
}

type ProjectCreateModel struct {
	Title       string `json:"title"`
	WorkspaceID uint   `json:"workspace_id"`
	OwnerID     *uint  `json:"owner_id,omitempty"`
}
