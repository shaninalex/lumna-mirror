// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name         string         `gorm:"size:255;not null"`
	Email        string         `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string         `gorm:"not null"`
	Active       bool           `gorm:"default:false"`
	CreatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// Implement IObject interface
func (s *UserModel) GetID() uuid.UUID       { return s.ID }
func (s *UserModel) SetID(id uuid.UUID)     { s.ID = id }
func (s *UserModel) GetUserId() uuid.UUID   { return s.ID }
func (s *UserModel) SetUserId(id uuid.UUID) { s.ID = id }

type UserRepository struct {
	Repository[*UserModel]
}

func NewUserRepository() *UserRepository {
	s := &UserRepository{}
	return s
}
