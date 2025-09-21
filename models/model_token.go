// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	Claims    string    `db:"claims"`
	Device    string    `db:"device"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}

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
