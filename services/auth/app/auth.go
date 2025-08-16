package app

import (
	"context"

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
	userRepository *database.UserRepository
}

func NewAuthApi() *AuthApi {
	return &AuthApi{
		userRepository: database.NewUserRepository(),
	}
}

func (s *AuthApi) HookRegister(ctx context.Context, data *domain.HooksKratosPayloadDTO) error {
	userId, err := uuid.Parse(data.UserID)
	if err != nil {
		return err
	}
	_, err = s.userRepository.Create(ctx, &database.UserModel{UserID: userId})
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthApi) HookVerify(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AuthApi) HookLogin(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
