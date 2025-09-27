// Copyright © 2025 Lumna. All rights reserved.

package models

import (
	"encoding/json"
)

// TaskStatus - task status.
type TaskStatus struct {
	ID        uint    `db:"id"`
	ProjectID uint    `db:"project_id"`
	Title     string  `db:"title"`
	Completed bool    `db:"complete"`
	ListIndex uint    `db:"list_index"`
	Config    *string `db:"config"`
}

// GetID - returns the id.
func (s *TaskStatus) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *TaskStatus) SetID(id uint) { s.ID = id }

// SaveConfig - saves the config.
func (s *TaskStatus) SaveConfig(cnf TaskStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	res := string(b)
	s.Config = &res
}

// GetConfig - returns the config.
func (s *TaskStatus) GetConfig() *TaskStatusConfig {
	if s.Config == nil {
		return NewTaskStatusConfig()
	}
	var config TaskStatusConfig
	err := json.Unmarshal([]byte(*s.Config), &config)
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
