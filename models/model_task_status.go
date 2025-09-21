// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// TaskStatus - task status.
type TaskStatus struct {
	ID          uuid.UUID `db:"id"`
	ProjectID   uuid.UUID `db:"project_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Complete    bool      `db:"complete"`
	Index       uint      `db:"index"`
	Config      string    `db:"config"`
}

// GetID - returns the id.
func (s *TaskStatus) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *TaskStatus) SetID(id uuid.UUID) { s.ID = id }

// SaveConfig - saves the config.
func (s *TaskStatus) SaveConfig(cnf TaskStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	s.Config = string(b)
}

// GetConfig - returns the config.
func (s *TaskStatus) GetConfig() *TaskStatusConfig {
	if s.Config == "" {
		return NewTaskStatusConfig()
	}
	var config TaskStatusConfig
	err := json.Unmarshal([]byte(s.Config), &config)
	if err != nil {
		return NewTaskStatusConfig()
	}
	return &config
}

// TaskStatusConfig - task status config.
type TaskStatusConfig struct {
	Color string `json:"color,omitempty"`
}

// NewTaskStatusConfig - new task status config.
func NewTaskStatusConfig() *TaskStatusConfig {
	return &TaskStatusConfig{
		Color: "default",
	}
}
