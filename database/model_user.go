package database

import (
	"time"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Settings  string
	Identity  *ory.Identity `gorm:"-"` // ignored by GORM
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// TODO: need to create user public code like @user123 . Save it in ory.Identity or in that model

	// no need to embedd this. Permissions can be changed during request and usermodel will can have old
	// permissions. Every time we need something from keto - ask it. Do not store!
	//Permissions any          `gorm:"-"` // Keto permissions data
}

// Implement IObject interface

func (s *User) GetID() uuid.UUID   { return s.ID }
func (s *User) SetID(id uuid.UUID) { s.ID = id }

type UserRepository struct {
	Repository[*User]
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

// User builder pattern code
type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	user := &User{}
	b := &UserBuilder{user: user}
	return b
}

func (b *UserBuilder) ID(iD uuid.UUID) *UserBuilder {
	b.user.ID = iD
	return b
}

func (b *UserBuilder) Settings(settings string) *UserBuilder {
	b.user.Settings = settings
	return b
}

func (b *UserBuilder) Identity(identity *ory.Identity) *UserBuilder {
	b.user.Identity = identity
	return b
}

func (b *UserBuilder) CreatedAt(createdAt time.Time) *UserBuilder {
	b.user.CreatedAt = createdAt
	return b
}

func (b *UserBuilder) UpdatedAt(updatedAt time.Time) *UserBuilder {
	b.user.UpdatedAt = updatedAt
	return b
}

func (b *UserBuilder) Build() *User {
	return b.user
}
