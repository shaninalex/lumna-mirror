package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID         uint     `gorm:"primaryKey"`
	IdentityID uint     `gorm:"not null;index"`
	Identity   Identity `gorm:"foreignKey:IdentityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Hashed refresh token
	Hash      string `gorm:"type:text;not null;uniqueIndex"`
	ClientID  string
	Scopes    string
	ExpiresAt time.Time `gorm:"not null;index"`
	Revoked   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null"`
}

func (u *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
