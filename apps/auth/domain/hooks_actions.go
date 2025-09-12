// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/apps/auth/dto"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/builders"
)

// AuthHookHandler - auth hook handler.
type AuthHookHandler interface {
	HookRegister(ctx context.Context, data *dto.HooksKratosPayloadDTO) error
	HookVerify(ctx context.Context) error
	HookLogin(ctx context.Context) error
}

var _ AuthHookHandler = &AuthHookAPI{}

// AuthHookAPI - auth hook api.
type AuthHookAPI struct {
}

// NewAuthHookAPI - new auth hook api.
func NewAuthHookAPI() *AuthHookAPI {
	return &AuthHookAPI{}
}

// HookRegister - hook register.
func (s *AuthHookAPI) HookRegister(ctx context.Context, data *dto.HooksKratosPayloadDTO) error {
	userID, err := uuid.Parse(data.UserID)
	if err != nil {
		return err
	}

	// TODO: move user settings creation into user service
	db := database.GetDB(ctx)
	userCode, err := database.GenerateUniqueUserCode(ctx, db, 5)
	if err != nil {
		log.Printf("unable to generate username code: %v\n", err)
		userCode = uuid.NewString()
	}
	user := &models.User{
		ID:   userID,
		Code: userCode,
	}
	tx := db.Create(&user)
	if tx.Error != nil {
		return apperrors.UserUnableToCreate
	}

	// move to separate user pipeline
	// For MVP we will just create organizations on user registration.
	// later when create organization and invitation steps will be completed - we remove that
	// It should be managed by organization microservice
	org := builders.NewOrganizationBuilder().
		Title(uuid.NewString()).
		UserID(user.GetID()).
		Build()
	tx = db.Create(&org)
	if tx.Error != nil {
		return tx.Error
	}
	// TODO: user.SetOrganization(org)
	db.Update("organization_id", &org.ID).Where("id = ?", userID)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// HookVerify - hook verify.
func (s *AuthHookAPI) HookVerify(ctx context.Context) error {
	// TODO: notify about Verify
	// TODO implement me
	panic("implement me")
}

// HookLogin - hook login.
func (s *AuthHookAPI) HookLogin(ctx context.Context) error {
	// TODO: notify about Login
	// TODO implement me
	panic("implement me")
}
