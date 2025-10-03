// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"time"
)

// Badge - task status.
type Badge struct {
	ID        uint      `db:"id"`
	ProjectID uint      `db:"project_id"`
	Title     string    `db:"title"`
	Config    string    `db:"complete"`
	CreatedAt time.Time `db:"created_at"`
}

// GetID - returns the id.
func (s *Badge) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *Badge) SetID(id uint) { s.ID = id }
