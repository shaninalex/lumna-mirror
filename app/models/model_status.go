package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type Status struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`

	Meta StatusMeta `gorm:"null" json:"meta"`

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
	Order    uint   `json:"order"`
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Expanded bool   `json:"expanded"`
}

func (j *StatusMeta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := StatusMeta{}
	err := json.Unmarshal(bytes, &result)
	*j = StatusMeta(result)
	return err
}

func (j StatusMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b).MarshalJSON()
}
