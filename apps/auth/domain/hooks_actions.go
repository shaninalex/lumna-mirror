// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/apps/auth/dto"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/models"
)

type AuthHookHandler interface {
	HookRegister(ctx context.Context, data *dto.HooksKratosPayloadDTO) error
	HookVerify(ctx context.Context) error
	HookLogin(ctx context.Context) error
}

var _ AuthHookHandler = &AuthHookApi{}

type AuthHookApi struct {
}

func NewAuthHookApi() *AuthHookApi {
	return &AuthHookApi{}
}

func (s *AuthHookApi) HookRegister(ctx context.Context, data *dto.HooksKratosPayloadDTO) error {
	userId, err := uuid.Parse(data.UserID)
	if err != nil {
		return err
	}

	// TODO: move to user service
	db := database.GetDB(ctx)
	userCode, err := database.GenerateUniqueUserCode(ctx, db, 5)
	if err != nil {
		log.Printf("unable to generate username code: %v\n", err)
		userCode = uuid.NewString()
	}
	user := &models.User{
		ID:   userId,
		Code: userCode,
	}
	tx := db.Create(&user)
	if tx.Error != nil {
		return err
	}
	return nil
}

func (s *AuthHookApi) HookVerify(ctx context.Context) error {
	// TODO: notify about Verify
	// TODO implement me
	panic("implement me")
}

func (s *AuthHookApi) HookLogin(ctx context.Context) error {
	// TODO: notify about Login
	// TODO implement me
	panic("implement me")
}
