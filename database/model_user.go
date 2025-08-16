// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package database

import (
	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
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

type UserRepository struct {
	Repository[*UserModel]
}

func NewUserRepository() *UserRepository {
	s := &UserRepository{}
	return s
}

// This method requires kratos client dependency which is quite bad for UserRepository
// Instead we need to create UserService which will provide userRepository methods
// and GetUser with kratos service method
//// GetUser main method to get fully defined user model
//func (s *UserRepository) GetUser(ctx context.Context, userID uuid.UUID) (*UserModel, error) {
//	var user UserModel
//	tx := s.DB.Where("user_id = ?", userID.String()).First(&user)
//	if tx.Error != nil {
//		return nil, tx.Error
//	}
//
//	identity, _, err := s.kratos.GetIdentity(ctx, userID.String())
//	if err != nil {
//		return nil, err
//	}
//	user.Identity = identity
//	return &user, nil
//}
