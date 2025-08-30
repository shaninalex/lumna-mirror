// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"net/http"

	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
)

// NewMockKratosService - new test kratos service.
func NewMockKratosService() *MockKratosService {
	return &MockKratosService{}
}

var _ kratos.IKratos = &MockKratosService{}

// MockKratosService - test kratos service.
type MockKratosService struct {
}

func (s *MockKratosService) Client() *ory.APIClient {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetLoginFlow(ctx context.Context, cookie, flowId string) (*ory.LoginFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetRegistrationFlow(ctx context.Context, cookie, flowId string) (*ory.RegistrationFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetErrorFlow(ctx context.Context, errorId string) (*ory.FlowError, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetVerificationFlow(ctx context.Context, cookie, flowId string) (*ory.VerificationFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) CreateLogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetSettingsFlow(ctx context.Context, cookie, flowId string) (*ory.SettingsFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetRecoveryFlow(ctx context.Context, cookie, flowId string) (*ory.RecoveryFlow, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockKratosService) GetSession(ctx context.Context, cookie string) (*ory.Session, *http.Response, error) {
	if session, ok := SessionStorage[cookie]; ok {
		return session, nil, nil
	}
	return nil, nil, apperrors.SessionNotFound
}

func (s *MockKratosService) GetIdentity(ctx context.Context, id string) (*ory.Identity, *http.Response, error) {
	for _, i := range IdentityStorage {
		if i.Id == id {
			return i, nil, nil
		}
	}
	return nil, nil, apperrors.UserIdentityNotFound
}

func (s *MockKratosService) UpdateIdentityTraits(ctx context.Context, id string, traits map[string]interface{}) (*ory.Identity, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}
