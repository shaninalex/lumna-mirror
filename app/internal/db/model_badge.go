package db

import (
	"time"
)

// Badge - task status.
type Badge struct {
	ID        int64     `db:"id"`
	ProjectID int64     `db:"project_id"`
	Title     string    `db:"title"`
	Config    string    `db:"complete"`
	CreatedAt time.Time `db:"created_at"`
}

// GetID - returns the id.
func (s *Badge) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *Badge) SetID(id int64) { s.ID = id }
