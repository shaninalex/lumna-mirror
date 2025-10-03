// Copyright © 2025 Lumna. All rights reserved.

package models

import (
	"encoding/json"
	"time"
)

type Badge struct {
	ID        uint
	ProjectID uint
	Title     string
	Config    *string
	CreatedAt time.Time
}

// SaveConfig - saves the config.
func (s *Badge) SaveConfig(cnf BadgeStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	res := string(b)
	s.Config = &res
}

// GetConfig - returns the config.
func (s *Badge) GetConfig() *BadgeStatusConfig {
	if s.Config == nil {
		return NewBadgeStatusConfig()
	}
	var config BadgeStatusConfig
	err := json.Unmarshal([]byte(*s.Config), &config)
	if err != nil {
		return NewBadgeStatusConfig()
	}
	return &config
}

// BadgeStatusConfig - badge config.
type BadgeStatusConfig struct {
	Color string `json:"color,omitempty"`
}

// NewBadgeStatusConfig - new task status config.
func NewBadgeStatusConfig() *BadgeStatusConfig {
	return &BadgeStatusConfig{
		Color: "default",
	}
}
