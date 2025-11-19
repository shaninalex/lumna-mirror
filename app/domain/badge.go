// Model Badge.
//
// Badge can be applied to entity and change it behaviour. For example tasks with blocked badge can't be
// complete until block is resolved. Or tasks with "not now" can't be started. Or for example "waiting for approval",
// "fix needed". You can imagine that situations.
//
// Badge is different from the "tag". Tag - help organize lists, badge - define special behaviour. It's not a "status",
// it's a "state". Please, do not be confused.

package domain

import (
	"encoding/json"
	"time"
)

type Badge struct {
	Id        int64       `json:"id"`
	ProjectID int64       `json:"project_id"`
	Title     string      `json:"title"`
	Config    BadgeConfig `json:"config"`
	CreatedAt time.Time   `json:"created_at"`
}

// BadgeConfig - badge config.
type BadgeConfig struct {
	Color string `json:"color,omitempty"`
}

// ToBadgeConfig - converts string to badge config.
func ToBadgeConfig(cnf string) BadgeConfig {
	var config BadgeConfig
	err := json.Unmarshal([]byte(cnf), &config)
	if err != nil {
		return NewBadgeStatusConfig()
	}
	return config
}

// NewBadgeStatusConfig - new task status config.
func NewBadgeStatusConfig() BadgeConfig {
	return BadgeConfig{
		Color: "default",
	}
}
