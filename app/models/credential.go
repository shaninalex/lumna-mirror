package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Credential struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	IdentityID uuid.UUID `gorm:"type:uuid;not null;index" json:"identity_id"`
	Identity   Identity  `gorm:"foreignKey:IdentityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Provider       string    `gorm:"type:text;not null;index:idx_provider_user,unique;index:idx_provider_email,unique"`
	ProviderUserID *string   `gorm:"type:text;index:idx_provider_user,unique" json:"provider_user_id,omitempty"`
	Email          *string   `gorm:"type:text;index:idx_provider_email,unique" json:"email,omitempty"`
	PasswordHash   *string   `gorm:"type:text" json:"-"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
}

func (u *Credential) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
