// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"
	"time"
)

type UserToken struct {
	ID        uint      `db:"id"`
	UserID    uint      `db:"user_id"`
	Claims    string    `db:"claims"`
	Device    string    `db:"device"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}

// GetID - returns the id.
func (s *UserToken) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *UserToken) SetID(id uint) { s.ID = id }

func (s *UserToken) SetClaims(claims map[string]any) {
	if b, err := json.Marshal(claims); err == nil {
		s.Claims = string(b)
	}
}

func (s *UserToken) GetClaims() map[string]any {
	claims := map[string]any{}
	if err := json.Unmarshal([]byte(s.Claims), &claims); err != nil {
		return claims
	}
	return claims
}
