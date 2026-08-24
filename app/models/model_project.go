package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type Project struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Title       string      `gorm:"not null" json:"title"`
	WorkspaceID uint        `gorm:"not null;index" json:"workspace_id"`
	OwnerID     *uint       `gorm:"null;index" json:"owner_id"`
	Meta        ProjectMeta `gorm:"meta" json:"meta"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
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

func (s *ProjectMeta) Scan(value interface{}) error {
	var data []byte

	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("unsupported type for ProjectMeta: %T", value)
	}

	var result ProjectMeta

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*s = result
	return nil
}

func (s ProjectMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return ProjectMeta{}, nil
	}
	return json.RawMessage(b).MarshalJSON()
}

type ProjectCreateModel struct {
	Title       string `json:"title"`
	WorkspaceID uint   `json:"workspace_id"`
	OwnerID     *uint  `json:"owner_id,omitempty"`
}
