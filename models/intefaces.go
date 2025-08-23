package models

import (
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type Model interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
}

type IUser interface {
	Model
	GetIdentity() *ory.Identity
	GetTraits() interface{}
}

type Ownable interface {
	Model
	GetOwnerID() uuid.UUID
	GetOwner() IUser
}

type Timestamped interface {
	Model
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
