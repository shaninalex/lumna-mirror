// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"time"

	"github.com/shaninalex/lumna/internal/db"
)

type UserManager interface {
	GetUser(ctx context.Context, userID uint) (*db.User, error)
	UpdateUserSettings(ctx context.Context, userID uint, settings *db.UserSettings) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser get user
func (s UserService) GetUser(ctx context.Context, userID uint) (*db.User, error) {
	user, err := db.UserGetByField(ctx, db.GetDb(ctx), "id", userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserSettings update user settings
func (s UserService) UpdateUserSettings(ctx context.Context, userID uint, settings *db.UserSettings) error {
	connection := db.GetDb(ctx)
	user, err := db.UserGetByField(ctx, connection, "id", userID)
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
