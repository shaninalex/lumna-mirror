// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserManager interface {
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUserSettings(ctx context.Context, userID uuid.UUID, settings *models.UserSettings) error
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user := models.User{ID: userID}
	tx := database.GetDB(ctx).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (s UserService) UpdateUserSettings(ctx context.Context, userID uuid.UUID, settings *models.UserSettings) error {
	user := models.User{ID: userID}
	db := database.GetDB(ctx)
	tx := db.First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	user.SetSettings(settings)
	tx = db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
