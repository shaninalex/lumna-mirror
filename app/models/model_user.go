package models

import (
	"time"
)

// and can be active/inactive
type AuthUser interface {
	Identifiable
	GetEmail() string
	SetEmail(string)
	IsActive() bool
	SetActive(bool)
	SetPassword(v string)
	GetPassword() string
}

type User struct {
	Id           uint      `db:"id"`
	Email        string    `db:"email"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PasswordHash string    `db:"password_hash"`
}

func (s *User) GetId() uint              { return s.Id }
func (s *User) SetId(v uint)             { s.Id = v }
func (s *User) GetEmail() string         { return s.Email }
func (s *User) SetEmail(v string)        { s.Email = v }
func (s *User) IsActive() bool           { return s.Active }
func (s *User) SetActive(v bool)         { s.Active = v }
func (s *User) SetPassword(v string)     { s.PasswordHash = v }
func (s *User) GetPassword() string      { return s.PasswordHash }
func (s *User) GetCreatedAt() time.Time  { return s.CreatedAt }
func (s *User) GetUpdatedAt() time.Time  { return s.UpdatedAt }
func (s *User) SetCreatedAt(v time.Time) { s.CreatedAt = v }
func (s *User) SetUpdatedAt(v time.Time) { s.UpdatedAt = v }
