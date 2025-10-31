// Copyright © 2025 Lumna. All rights reserved.

package db

// TaskStatus - task status.
type TaskStatus struct {
	ID        int64   `db:"id"`
	ProjectID int64   `db:"project_id"`
	Title     string  `db:"title"`
	ListIndex int64   `db:"list_index"`
	Config    *string `db:"config"`
}

// GetID - returns the id.
func (s *TaskStatus) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *TaskStatus) SetID(id int64) { s.ID = id }
