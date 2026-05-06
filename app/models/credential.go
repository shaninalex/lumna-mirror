package models

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	IdentityID     uint      `gorm:"not null;index" json:"identity_id"`
	Provider       string    `gorm:"type:text;not null;index:idx_provider_user,unique;index:idx_provider_email,unique"`
	ProviderUserID *string   `gorm:"type:text;index:idx_provider_user,unique" json:"provider_user_id,omitempty"`
	Email          *string   `gorm:"type:text;index:idx_provider_email,unique" json:"email,omitempty"`
	PasswordHash   *string   `gorm:"type:text" json:"-"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
}

func (u *Credential) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
