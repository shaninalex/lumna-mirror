package domain

import (
	"encoding/json"
	"log"
	"time"
)

type (

	// UserState defines user's state
	UserState string

	// User user model
	User struct {
		Id           int64
		Email        string
		Settings     string
		Active       bool
		State        UserState
		Code         string
		PasswordHash string `json:"-"`

		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

// GetID - returns the id.
func (s *User) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *User) SetID(id int64) { s.Id = id }

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

const (
	UserStatePending UserState = "pending"
	UserStateActive  UserState = "active"
	UserStateDeleted UserState = "deleted"
	UserStateBanned  UserState = "banned"
)
