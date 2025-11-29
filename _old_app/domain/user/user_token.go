package user

import (
	"time"
)

type UserToken struct {
	Id               int64      `db:"id"`
	UserID           int64      `db:"user_id"`
	Device           string     `db:"device"`
	RefreshToken     string     `db:"refresh_token"`
	RefreshExpiresAt time.Time  `db:"refresh_expires_at"`
	Revoked          bool       `db:"revoked"`
	RevokedAt        *time.Time `db:"revoked_at"`
	CreatedAt        time.Time  `db:"created_at"`
}

// GetID - returns the id.
func (s *UserToken) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *UserToken) SetID(id int64) { s.Id = id }

// IsExpired - check is token expired
func (s *UserToken) IsExpired() bool {
	return s.RefreshExpiresAt.Before(time.Now())
}
