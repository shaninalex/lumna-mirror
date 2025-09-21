// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// UserBuilder builder pattern code
type UserBuilder struct {
	user *models.User
}

// NewUserBuilder - new user builder.
func NewUserBuilder() *UserBuilder {
	user := &models.User{}
	b := &UserBuilder{user: user}
	return b
}

// ID - id.
func (b *UserBuilder) ID(iD uuid.UUID) *UserBuilder {
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
func (b *UserBuilder) Build() *models.User {
	return b.user
}
