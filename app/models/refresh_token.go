package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// FK → identities(id)
	IdentityID uuid.UUID `gorm:"type:uuid;not null;index" json:"identity_id"`
	Identity   Identity  `gorm:"foreignKey:IdentityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Hashed refresh token (never store raw token)
	TokenHash string `gorm:"type:text;not null;uniqueIndex" json:"-"`

	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	Revoked   bool      `gorm:"not null;default:false" json:"revoked"`

	UserAgent *string `gorm:"type:text" json:"user_agent,omitempty"`
	IpAddress *string `gorm:"type:text" json:"ip_address,omitempty"`

	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (u *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
