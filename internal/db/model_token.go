// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"time"
)

type UserToken struct {
	ID               uint       `db:"id"`
	UserID           uint       `db:"user_id"`
	Device           string     `db:"device"`
	RefreshToken     string     `db:"refresh_token"`
	RefreshExpiresAt time.Time  `db:"refresh_expires_at"`
	Revoked          bool       `db:"revoked"`
	RevokedAt        *time.Time `db:"revoked_at"`
	CreatedAt        time.Time  `db:"created_at"`
}

// GetID - returns the id.
func (s *UserToken) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *UserToken) SetID(id uint) { s.ID = id }

// IsExpired - check is token expired
func (s *UserToken) IsExpired() bool {
	return s.RefreshExpiresAt.Before(time.Now())
}
