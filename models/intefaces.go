// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

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
	GetSettings() *UserSettings
	SetSettings(*UserSettings)
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
}

// Auditable get id of an entity created
type Auditable interface {
	GetCreatedBy() uuid.UUID
}

// Coded describe an entity has and use codes
type Coded interface {
	SetCode(string)
	GetCode() string
}
