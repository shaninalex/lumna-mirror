package models

import (
	"time"
)

// SessionModel represents a session for a user
type SessionModel struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Data      string    `db:"data"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}
