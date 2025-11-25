package project

import (
	"time"
)

type Project struct {
	Id        int64
	Title     string
	Code      string
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetID - returns the id.
func (s *Project) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Project) SetID(id int64) { s.Id = id }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() int64 { return s.UserId }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() int64 { return s.GetOwnerID() }
