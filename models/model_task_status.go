// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type TaskStatus struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ProjectID uuid.UUID
	Project   *Project `gorm:"foreignKey:ProjectID;references:ID"`

	Tasks []*Task `gorm:"foreignKey:TaskStatusID;references:ID"`

	Title       string `gorm:"uniqueIndex"`
	Description string
	Complete    bool
	Index       uint
	Config      string
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

type TaskStatusConfig struct {
	Color string `json:"color,omitempty"`
}

func NewTaskStatusConfig() *TaskStatusConfig {
	return &TaskStatusConfig{
		Color: "default",
	}
}
