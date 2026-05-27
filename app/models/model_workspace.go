package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/shaninalex/lumna/app/pkg/observer"
)

var (
	EventWorkspaceCreated observer.Event = "WORKSPACE_CREATED"
)

type Workspace struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Slug       string     `gorm:"slug" json:"slug"`
	Title      string     `json:"title"`
	Active     bool       `json:"active"`
	OwnerEmail string     `json:"owner_email"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type WorkspaceCreateModel struct {
	Title string `json:"title" binding:"required"`
}

func (a WorkspaceCreateModel) Validate() error {
	return validation.ValidateStruct(&a,
		// Title cannot be empty, and the length must between 3 and 50
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
	)
}
