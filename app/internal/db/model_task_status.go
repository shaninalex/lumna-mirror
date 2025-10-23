// Copyright © 2025 Lumna. All rights reserved.

package db

// TaskStatus - task status.
type TaskStatus struct {
	ID        uint    `db:"id"`
	ProjectID uint    `db:"project_id"`
	Title     string  `db:"title"`
	ListIndex uint    `db:"list_index"`
	Config    *string `db:"config"`
}

// GetID - returns the id.
func (s *TaskStatus) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *TaskStatus) SetID(id uint) { s.ID = id }
