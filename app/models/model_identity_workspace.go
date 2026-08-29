package models

import (
	"time"
)

type IdentityWorkspace struct {
	IdentityId  uint `gorm:"primaryKey"`
	WorkspaceId uint `gorm:"primaryKey"`

	// CreatedAt - when user was joined to the workspace
	CreatedAt time.Time
}
