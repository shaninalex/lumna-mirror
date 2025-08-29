package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID *uuid.UUID
	Organization   *Organization `gorm:"foreignKey:OrganizationID;references:ID"`

	Settings string
	Identity *ory.Identity `gorm:"-"` // ignored by GORM
	Code     string        `gorm:"uniqueIndex"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// TODO: need to create user public code like @user123 . Save it in ory.Identity or in that model

	// no need to embedd this. Permissions can be changed during request and usermodel will can have old
	// permissions. Every time we need something from keto - ask it. Do not store!
	//Permissions any          `gorm:"-"` // Keto permissions data
}

func (s *User) GetID() uuid.UUID   { return s.ID }
func (s *User) SetID(id uuid.UUID) { s.ID = id }
func (s *User) GetTraits() any     { return s.GetIdentity().GetTraits() }
func (s *User) GetIdentity() *ory.Identity {
	if s.Identity == nil {
		panic(fmt.Errorf("identity not set"))
	}
	return s.Identity
}
func (s *User) IsActive() bool          { return s.GetIdentity().GetState() == "active" }
func (s *User) GetCreatedAt() time.Time { return s.CreatedAt }
func (s *User) GetUpdatedAt() time.Time { return s.UpdatedAt }
func (s *User) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}
func (s *User) IsDeleted() bool { return s.DeletedAt.Valid }

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
