package models

import (
	"time"

	"github.com/google/uuid"
)

// Identifiable each database model has ID and can set/get its
type Identifiable interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
}

// AuthUser describe an authenticated model interface that should has Traits ( email, name, code etc )
// and can be active/inactive
type AuthUser interface {
	Identifiable
	GetTraits() any
	IsActive() bool
}

// Ownable has an owner and can validate other
type Ownable interface {
	Identifiable
	GetOwnerID() uuid.UUID
	GetOwner() AuthUser
	IsOwner(entity AuthUser) bool
}

// Timestamped is a model that has timestamped fields
type Timestamped interface {
	Identifiable
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time // optional for soft-delete
	IsDeleted() bool
}

type Auditable interface {
	GetCreatedBy() uuid.UUID
}
