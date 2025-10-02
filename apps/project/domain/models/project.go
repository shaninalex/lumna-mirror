// Copyright © 2025 Lumna. All rights reserved.

package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID        uint
	Title     string
	Code      string
	Statuses  []Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID        uint
	Title     string
	Idx       uint
	Completed bool
	ListIndex uint
	Config    *string
}

// SaveConfig - saves the config.
func (s *Status) SaveConfig(cnf TaskStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	res := string(b)
	s.Config = &res
}

// GetConfig - returns the config.
func (s *Status) GetConfig() *TaskStatusConfig {
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
