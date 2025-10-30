// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"time"

	"github.com/shaninalex/lumna/app/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type (

	// UserState defines user's state
	UserState string

	// User user model
	User struct {
		ID       uint
		Email    string
		Settings string
		Active   bool
		State    UserState
		Code     string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

const (
	UserStatePending UserState = "pending"
	UserStateActive  UserState = "active"
	UserStateDeleted UserState = "deleted"
	UserStateBanned  UserState = "banned"
)

type UserManager interface {
	GetUser(ctx context.Context, userID uint) (*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	UpdateUserSettings(ctx context.Context, userID uint, settings *db.UserSettings) error
	SetPassword(ctx context.Context, userId uint, rawPwd string) error
	CheckPassword(ctx context.Context, userId uint, password string) (bool, error)
	CreateUser(ctx context.Context, email string) *User
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser get user
func (s UserService) GetUser(ctx context.Context, userId uint) (*db.User, error) {
	user, err := db.UserGetByField(ctx, db.GetDb(ctx), "id", userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail get user by email
func (s UserService) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := db.UserGetByField(ctx, db.GetDb(ctx), "email", email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserSettings update user settings
func (s UserService) UpdateUserSettings(ctx context.Context, userId uint, settings *db.UserSettings) error {
	connection := db.GetDb(ctx)
	user, err := db.UserGetByField(ctx, connection, "id", userId)
	if err != nil {
		return err
	}
	user.SetSettings(settings)
	user.UpdatedAt = time.Now()
	if err = db.UserUpdate(ctx, connection, user); err != nil {
		return err
	}
	return nil
}

// SetPassword set password
func (s UserService) SetPassword(ctx context.Context, userId uint, rawPwd string) error {
	panic("not implemented")
}

// CheckPassword check password
func (s UserService) CheckPassword(ctx context.Context, userId uint, password string) error {
	user, err := db.UserGetByField(ctx, db.GetDb(ctx), "id", userId)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return err
	}
	return nil
}

// CreateUser create user
func (s UserService) CreateUser(ctx context.Context, email, rawPassword string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &db.User{
		Email:        email,
		PasswordHash: string(hash),
	}

	connection := db.GetDb(ctx)
	user, err = db.UserSave(ctx, connection, user)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        user.ID,
		Email:     user.Email,
		Settings:  user.Settings,
		Active:    user.Active,
		State:     UserState(user.State),
		Code:      user.Code,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
