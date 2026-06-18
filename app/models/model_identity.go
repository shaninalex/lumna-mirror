package models

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type IdentityID uint

type Identity struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	FullName  string     `gorm:"not null" json:"full_name"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Active    bool       `json:"active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *Identity) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}

func (u *Identity) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = utils.Pointer(time.Now())
	return nil
}

func (s *Identity) String() string {
	return fmt.Sprintf("Identity id=%d email=%s", s.ID, s.Email)
}
