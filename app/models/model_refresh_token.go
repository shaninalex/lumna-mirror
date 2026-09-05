package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID         int    `gorm:"primaryKey"`
	IdentityID int    `gorm:"not null;index"`
	Hash       string `gorm:"type:text;not null;uniqueIndex"`
	ClientID   string
	Scopes     string
	ExpiresAt  time.Time `gorm:"not null;index"`
	Revoked    bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"not null"`
}

func (u *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
