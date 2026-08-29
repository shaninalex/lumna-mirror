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

type Column struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`

	Meta ColumnMeta `gorm:"null" json:"meta"`

	BoardId uint `gorm:"not null;index" json:"board_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *Column) BeforeCreate(tx *gorm.DB) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	return nil
}

func (s *Column) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Column) String() string {
	return fmt.Sprintf("Column id=%d title=%s", s.ID, s.Title)
}

type ColumnMeta struct {
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Expanded bool   `json:"expanded"`
}

func (j *ColumnMeta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := ColumnMeta{}
	err := json.Unmarshal(bytes, &result)
	*j = ColumnMeta(result)
	return err
}

func (j ColumnMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b).MarshalJSON()
}
