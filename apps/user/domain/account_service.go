package domain

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
)

type AccountManager interface {
	LogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, error)
	GetAccountSettingsForm(ctx context.Context, cookie string, flowID uuid.UUID) (*ory.SettingsFlow, error)
	CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, error)
}

func NewAccountService(kratosService kratos.IKratos) *AccountService {
	return &AccountService{
		kratosService: kratosService,
	}
}

type AccountService struct {
	kratosService kratos.IKratos
}

func (s AccountService) LogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, error) {
	flow, _, err := s.kratosService.CreateLogoutFlow(ctx, cookie)
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (s AccountService) GetAccountSettingsForm(ctx context.Context, cookie string, flowID uuid.UUID) (*ory.SettingsFlow, error) {
	flow, _, err := s.kratosService.GetSettingsFlow(ctx, cookie, flowID.String())
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (s AccountService) CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, error) {
	flow, _, err := s.kratosService.CreateSettingsFlow(ctx, cookie)
	if err != nil {
		return nil, err
	}
	return flow, nil
}
