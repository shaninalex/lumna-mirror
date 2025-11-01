package builders

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

// UserBuilder builder pattern code
type UserBuilder struct {
	user *db.User
}

// NewUserBuilder - new user builder.
func NewUserBuilder() *UserBuilder {
	user := &db.User{}
	b := &UserBuilder{user: user}
	return b
}

// ID - id.
func (b *UserBuilder) ID(iD int64) *UserBuilder {
	b.user.ID = iD
	return b
}

// Settings - settings.
func (b *UserBuilder) Settings(settings string) *UserBuilder {
	b.user.Settings = settings
	return b
}

// Code - code.
func (b *UserBuilder) Code(code string) *UserBuilder {
	b.user.Code = code
	return b
}

// CreatedAt - created at.
func (b *UserBuilder) CreatedAt(createdAt time.Time) *UserBuilder {
	b.user.CreatedAt = createdAt
	return b
}

// UpdatedAt - updated at.
func (b *UserBuilder) UpdatedAt(updatedAt time.Time) *UserBuilder {
	b.user.UpdatedAt = updatedAt
	return b
}

// Build - builds the value.
func (b *UserBuilder) Build() *db.User {
	return b.user
}
