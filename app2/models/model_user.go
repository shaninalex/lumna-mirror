package models

import (
	"time"
)

// AuthUser describe user account
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

type UserToken struct {
	Id     uint `db:"id"`
	UserId uint `db:"user_id"`
	// Device           string     `db:"device"`
	RefreshToken     string     `db:"refresh_token"`
	RefreshExpiresAt time.Time  `db:"refresh_expires_at"`
	Revoked          bool       `db:"revoked"`
	RevokedAt        *time.Time `db:"revoked_at"`
	CreatedAt        time.Time  `db:"created_at"`
}

// GetI - returns the id.
func (s *UserToken) GetId() uint { return s.Id }

// SetID - sets the id.
func (s *UserToken) SetId(id uint) { s.Id = id }

// IsExpired - check is token expired
func (s *UserToken) IsExpired() bool {
	return s.RefreshExpiresAt.Before(time.Now())
}

// IsRevoked - check is token revoked
func (s *UserToken) IsRevoked() bool {
	return s.Revoked
}
