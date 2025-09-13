// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gorm.io/gorm"
)

// User - user.
type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID *uuid.UUID
	Organization   *Organization `gorm:"foreignKey:OrganizationID;references:ID"`

	Settings string
	Identity *ory.Identity `gorm:"-"` // ignored by GORM
	Code     string        `gorm:"uniqueIndex"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// TODO: need to create user public code like @user123 . Save it in ory.Identity or in that model

	// no need to embedd this. Permissions can be changed during request and usermodel will can have old
	// permissions. Every time we need something from keto - ask it. Do not store!
	//Permissions any          `gorm:"-"` // Keto permissions data
}

// GetID - returns the id.
func (s *User) GetID() uuid.UUID { return s.ID }

// SetID - sets the id.
func (s *User) SetID(id uuid.UUID) { s.ID = id }

// GetTraits - returns the traits.
func (s *User) GetTraits() any { return s.GetIdentity().GetTraits() }

// GetIdentity - returns the identity.
func (s *User) GetIdentity() *ory.Identity {
	if s.Identity == nil {
		panic(fmt.Errorf("identity not set"))
	}
	return s.Identity
}

// IsActive - checks if it is active.
func (s *User) IsActive() bool { return s.GetIdentity().GetState() == "active" }

// GetCreatedAt - returns the created at.
func (s *User) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *User) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetDeletedAt - returns the deleted at.
func (s *User) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}

// IsDeleted - checks if it is deleted.
func (s *User) IsDeleted() bool { return s.DeletedAt.Valid }

// SetCode - sets the code.
func (s *User) SetCode(code string) { s.Code = code }

// GetCode - returns the code.
func (s *User) GetCode() string { return s.Code }

func (s *User) GetSettings() *UserSettings {
	var settings UserSettings
	err := json.Unmarshal([]byte(s.Settings), &settings)
	if err != nil {
		panic(err)
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

// TraitsName - traits name.
type TraitsName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// UserTraits - user traits.
type UserTraits struct {
	Email string     `json:"email"`
	Name  TraitsName `json:"name"`
}

type Language string

const (
	LanguageEn Language = "en"
	LanguageUa Language = "ua"
	LanguageDe Language = "de"
)

type Theme string

const (
	ThemeLight  Theme = "light"
	ThemeDark   Theme = "dark"
	ThemeDevice Theme = "device"
)

type UserSettings struct {
	Theme        Theme    `json:"theme" validate:"required"`
	Language     Language `json:"language" validate:"required"`
	Timezone     string   `json:"timezone" validate:"required"`
	DateFormat   string   `json:"date_format" validate:"required"`
	TimeFormat   string   `json:"time_format" validate:"required"`
	WeekStartDay string   `json:"week_start_day" validate:"required"`
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
	WeekStartDay: "1",
}
