// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package models

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	UserID   uuid.UUID
	Settings string
	Identity *ory.Identity `gorm:"-"` // Kratos identity information

	// no need to embedd this. Permissions can be changed during request and usermodel will can have old
	// permissions. Every time we need something from keto - ask it. Do not store!
	//Permissions any          `gorm:"-"` // Keto permissions data
}

// Implement IObject interface
func (s *UserModel) GetID() uint   { return s.ID }
func (s *UserModel) SetID(id uint) { s.ID = id }

// func (s *UserModel) AfterFind(tx *gorm.DB) (err error) {
// 	// set user data from kratos
// 	// set permission data from keto
// 	return
// }

type UserRepository struct {
	Repository[*UserModel]
	kratos kratos.IKratos
}

func NewUserRepository(k kratos.IKratos) *UserRepository {
	s := &UserRepository{
		kratos: k,
	}
	return s
}

// GetUser main method to get fully defined user model
func (s *UserRepository) GetUser(ctx context.Context, userID uuid.UUID) (*UserModel, error) {
	var user UserModel
	tx := s.DB.Where("user_id = ?", userID.String()).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	identity, err := s.kratos.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.Identity = identity
	return &user, nil
}
