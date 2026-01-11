package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Identity struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `gorm:"unique" json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Identity) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}

func (u *Identity) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}
