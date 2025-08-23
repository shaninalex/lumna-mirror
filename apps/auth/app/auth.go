package app

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/domain"
)

type IAuthApi interface {
	HookRegister(ctx context.Context, data *domain.HooksKratosPayloadDTO) error
	HookVerify(ctx context.Context) error
	HookLogin(ctx context.Context) error
}

var _ IAuthApi = &AuthApi{}

type AuthApi struct {
}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

func (s *AuthApi) HookRegister(ctx context.Context, data *domain.HooksKratosPayloadDTO) error {
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
	user := &database.User{
		ID:   userId,
		Code: userCode,
	}
	tx := db.Create(&user)
	if tx.Error != nil {
		return err
	}
	return nil
}

func (s *AuthApi) HookVerify(ctx context.Context) error {
	// TODO: notify about Verify
	// TODO implement me
	panic("implement me")
}

func (s *AuthApi) HookLogin(ctx context.Context) error {
	// TODO: notify about Login
	// TODO implement me
	panic("implement me")
}
