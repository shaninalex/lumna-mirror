package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

type UserManager interface {
	GetUser(ctx context.Context, userID int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUserSettings(ctx context.Context, userID int64, settings *UserSettings) error
	SetPassword(ctx context.Context, userId int64, rawPwd string) error
	CheckPassword(ctx context.Context, userId int64, rawPassword string) error
	CreateUser(ctx context.Context, email string, rawPassword string) (*User, error)
}

var _ UserManager = &UserService{}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser get user
func (s UserService) GetUser(ctx context.Context, userId int64) (*User, error) {
	user, err := UserGetByField(ctx, db.GetDb(ctx), "id", userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail get user by email
func (s UserService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user, err := UserGetByField(ctx, db.GetDb(ctx), "email", email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserSettings update user settings
func (s UserService) UpdateUserSettings(ctx context.Context, userId int64, settings *UserSettings) error {
	connection := db.GetDb(ctx)
	user, err := UserGetByField(ctx, connection, "id", userId)
	if err != nil {
		return err
	}
	user.SetSettings(settings)
	user.UpdatedAt = time.Now()
	if err = UserUpdate(ctx, connection, user); err != nil {
		return err
	}
	return nil
}

// SetPassword set password
func (s UserService) SetPassword(ctx context.Context, userId int64, rawPwd string) error {
	panic("not implemented")
}

// CheckPassword check password
func (s UserService) CheckPassword(ctx context.Context, userId int64, password string) error {
	user, err := UserGetByField(ctx, db.GetDb(ctx), "id", userId)
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

	user := &User{
		Email:        email,
		PasswordHash: string(hash),
	}

	connection := db.GetDb(ctx)
	user, err = UserSave(ctx, connection, user)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:        user.Id,
		Email:     user.Email,
		Settings:  user.Settings,
		Active:    user.Active,
		State:     UserState(user.State),
		Code:      user.Code,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
