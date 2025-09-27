// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"time"
)

type UserToken struct {
	ID               uint       `db:"id"`
	Jti              string     `db:"jti"`
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
