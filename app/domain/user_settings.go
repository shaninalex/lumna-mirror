package domain

import (
	"encoding/json"
)

type (
	UserSettings struct {
		Theme        string `json:"theme" validate:"required"`
		Language     string `json:"language" validate:"required"`
		Timezone     string `json:"timezone" validate:"required"`
		DateFormat   string `json:"date_format" validate:"required"`
		TimeFormat   string `json:"time_format" validate:"required"`
		WeekStartDay int64  `json:"week_start_day" validate:"required"`
	}
)

func (s *UserSettings) ToString() string {
	if b, err := json.Marshal(s); err == nil {
		return string(b)
	}
	return "{}"
}

var DefaultUserSettings = UserSettings{
	Theme:        "light",
	Language:     "en",
	Timezone:     "UTC",
	DateFormat:   "02.01.2006",
	TimeFormat:   "15:04",
	WeekStartDay: 1,
}
