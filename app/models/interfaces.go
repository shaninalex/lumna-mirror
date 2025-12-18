package models

import (
	"time"
)

// Identifiable each database model has Id and can set/get its
type Identifiable interface {
	GetId() uint
	SetId(uint)
}

// Ownable has an owner and can validate other
type Ownable interface {
	Identifiable
	GetOwnerId() uint
	IsOwner(entity AuthUser) bool
}

// Timestamped is a model that has timestamped fields
type Timestamped interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetCreatedAt(v time.Time)
	SetUpdatedAt(v time.Time)
}
