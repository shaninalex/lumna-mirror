package models

import "time"

type WorkspaceModel struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Title      string     `json:"title"`
	Active     bool       `json:"active"`
	OwnerEmail string     `json:"owner_email"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
