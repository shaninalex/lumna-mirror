// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

type UserManager interface {
	GetUser(ctx context.Context, userID uint) (*models.User, error)
	UpdateUserSettings(ctx context.Context, userID uint, settings *models.UserSettings) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser get user
func (s UserService) GetUser(ctx context.Context, userID uint) (*models.User, error) {
	user, err := repositories.UserGetByField(ctx, database.GetDb(ctx), "id", userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserSettings update user settings
func (s UserService) UpdateUserSettings(ctx context.Context, userID uint, settings *models.UserSettings) error {
	db := database.GetDb(ctx)
	user, err := repositories.UserGetByField(ctx, db, "id", userID)
	if err != nil {
		return err
	}
	user.SetSettings(settings)
	user.UpdatedAt = time.Now()
	if err = repositories.UserUpdate(ctx, db, user); err != nil {
		return err
	}
	return nil
}
