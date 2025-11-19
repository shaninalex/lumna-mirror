package domain

import (
	"time"
)

type Comment struct {
	Id         int64     `json:"id" db:"id"`
	EntityId   int64     `json:"entity_id" db:"entity_id"`
	EntityType string    `json:"entity_type" db:"entity_type"`
	UserId     int64     `json:"user_id" db:"project_id"`
	Content    string    `json:"content" db:"title"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// GetID - returns the id.
func (s *Comment) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Comment) SetID(id int64) { s.Id = id }

// GetOwnerID - returns the owner id.
func (s *Comment) GetOwnerID() int64 { return s.UserId }

// IsOwner - checks if it is owner.
func (s *Comment) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Comment) GetCreatedAt() time.Time { return s.CreatedAt }

// GetCreatedBy - returns the created by.
func (s *Comment) GetCreatedBy() int64 { return s.UserId }
