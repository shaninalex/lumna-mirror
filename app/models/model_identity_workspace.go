package models

import (
	"time"
)

type IdentityWorkspace struct {
	IdentityId  int `gorm:"primaryKey"`
	WorkspaceId int `gorm:"primaryKey"`

	// CreatedAt - when user was joined to the workspace
	CreatedAt time.Time
}
