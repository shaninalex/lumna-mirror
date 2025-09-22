// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"
	"log"
	"time"
)

type UserState string

const (
	UserStatePending UserState = "pending"
	UserStateActive  UserState = "active"
	UserStateDeleted UserState = "deleted"
	UserStateBanned  UserState = "banned"
)

// User - user.
type User struct {
	ID           uint      `db:"id"`
	Email        string    `db:"email"`
	Settings     string    `db:"settings"`
	Active       bool      `db:"active"`
	State        UserState `db:"state"`
	Code         string    `db:"code"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// GetID - returns the id.
func (s *User) GetID() uint { return s.ID }

// SetID - sets the id.
func (s *User) SetID(id uint) { s.ID = id }

// IsActive - checks if it is active.
func (s *User) IsActive() bool { return s.Active }

// GetCreatedAt - returns the created at.
func (s *User) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *User) GetUpdatedAt() time.Time { return s.UpdatedAt }

// SetCode - sets the code.
func (s *User) SetCode(code string) { s.Code = code }

// GetCode - returns the code.
func (s *User) GetCode() string { return s.Code }

func (s *User) GetSettings() *UserSettings {
	var settings UserSettings
	err := json.Unmarshal([]byte(s.Settings), &settings)
	if err != nil {
		log.Println("User.GetSettings. Error:", err)
		return &DefaultUserSettings
	}
	return &settings
}

func (s *User) SetSettings(settings *UserSettings) {
	b, err := json.Marshal(&settings)
	if err != nil {
		panic(err)
	}
	s.Settings = string(b)
}

type UserSettings struct {
	Theme        string `json:"theme" validate:"required"`
	Language     string `json:"language" validate:"required"`
	Timezone     string `json:"timezone" validate:"required"`
	DateFormat   string `json:"date_format" validate:"required"`
	TimeFormat   string `json:"time_format" validate:"required"`
	WeekStartDay uint   `json:"week_start_day" validate:"required"`
}

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
