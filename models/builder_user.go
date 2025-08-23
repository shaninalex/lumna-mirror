package models

import (
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

// UserBuilder builder pattern code
type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	user := &User{}
	b := &UserBuilder{user: user}
	return b
}

func (b *UserBuilder) ID(iD uuid.UUID) *UserBuilder {
	b.user.ID = iD
	return b
}

func (b *UserBuilder) Settings(settings string) *UserBuilder {
	b.user.Settings = settings
	return b
}

func (b *UserBuilder) Code(code string) *UserBuilder {
	b.user.Code = code
	return b
}

func (b *UserBuilder) Identity(identity *ory.Identity) *UserBuilder {
	b.user.Identity = identity
	return b
}

func (b *UserBuilder) CreatedAt(createdAt time.Time) *UserBuilder {
	b.user.CreatedAt = createdAt
	return b
}

func (b *UserBuilder) UpdatedAt(updatedAt time.Time) *UserBuilder {
	b.user.UpdatedAt = updatedAt
	return b
}

func (b *UserBuilder) Build() *User {
	return b.user
}
