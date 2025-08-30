// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package kratos

import (
	"context"
	"net/http"

	ory "github.com/ory/kratos-client-go"
)

// IKratos - i kratos.
type IKratos interface {
	Client() *ory.APIClient
	GetLoginFlow(ctx context.Context, cookie, flowId string) (*ory.LoginFlow, *http.Response, error)
	GetRegistrationFlow(ctx context.Context, cookie, flowId string) (*ory.RegistrationFlow, *http.Response, error)
	GetErrorFlow(ctx context.Context, errorId string) (*ory.FlowError, *http.Response, error)
	GetVerificationFlow(ctx context.Context, cookie, flowId string) (*ory.VerificationFlow, *http.Response, error)
	CreateLogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, *http.Response, error)
	CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, *http.Response, error)
	GetSettingsFlow(ctx context.Context, cookie, flowId string) (*ory.SettingsFlow, *http.Response, error)
	GetRecoveryFlow(ctx context.Context, cookie, flowId string) (*ory.RecoveryFlow, *http.Response, error)
	GetSession(ctx context.Context, cookie string) (*ory.Session, *http.Response, error)
	GetIdentity(ctx context.Context, id string) (*ory.Identity, *http.Response, error)
	UpdateIdentityTraits(ctx context.Context, id string, traits map[string]interface{}) (*ory.Identity, *http.Response, error)
}

// KratosService - kratos service.
type KratosService struct {
	client *ory.APIClient
}

// NewKratosService - new kratos service.
func NewKratosService(url string) *KratosService {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: url,
		},
	}
	return &KratosService{
		client: ory.NewAPIClient(configuration),
	}
}

// Client return ory kratos api client itself
func (s *KratosService) Client() *ory.APIClient {
	return s.client
}

// GetLoginFlow is the method that return created login flow based on user
// cookies and flow id
func (s *KratosService) GetLoginFlow(ctx context.Context, cookie, flowId string) (*ory.LoginFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetLoginFlow(ctx).
		Cookie(cookie).
		Id(flowId).
		Execute()
}

// GetRegistrationFlow is the method that return registration flow based on user
// cookies and flow id
func (s *KratosService) GetRegistrationFlow(ctx context.Context, cookie, flowId string) (*ory.RegistrationFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetRegistrationFlow(ctx).
		Cookie(cookie).
		Id(flowId).
		Execute()
}

// GetErrorFlow is the method that return error flow based on user
// cookies and flow id. ory/kratos store errors in database and user can
// access them via error flow id.
func (s *KratosService) GetErrorFlow(ctx context.Context, errorId string) (*ory.FlowError, *http.Response, error) {
	return s.client.FrontendAPI.
		GetFlowError(ctx).
		Id(errorId).
		Execute()
}

// GetVerificationFlow is the method that return verification flow based on user
// cookies and flow id. After registration user can be redirected to the
// verification page where this form should be rendered.
func (s *KratosService) GetVerificationFlow(ctx context.Context, cookie, flowId string) (*ory.VerificationFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetVerificationFlow(ctx).
		Cookie(cookie).
		Id(flowId).
		Execute()
}

// CreateLogoutFlow is the method that create browser logout flow. This flow
// return logout token and logout url. By following logout id or calling api
// call with logout token - kratos will remove active session from db and user
// will be logged out
func (s *KratosService) CreateLogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		CreateBrowserLogoutFlow(ctx).
		Cookie(cookie).
		Execute()
}

// CreateSettingsFlow is the method that create user settings flow. This method
// basically says to the system that this user ( by cookie ) want to get his
// settings form to change something about himself.
func (s *KratosService) CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		CreateBrowserSettingsFlow(ctx).
		Cookie(cookie).
		Execute()
}

// GetSettingsFlow is the method that return settings flow to user. This object
// contain all editable user profile fields that kratos user have access to.
func (s *KratosService) GetSettingsFlow(ctx context.Context, cookie, flowId string) (*ory.SettingsFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetSettingsFlow(ctx).
		Cookie(cookie).
		Id(flowId).
		Execute()
}

// GetRecoveryFlow is the method that return account recovery flow
func (s *KratosService) GetRecoveryFlow(ctx context.Context, cookie, flowId string) (*ory.RecoveryFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetRecoveryFlow(ctx).
		Cookie(cookie).
		Id(flowId).
		Execute()
}

// GetSession is the method that return user session information
func (s *KratosService) GetSession(ctx context.Context, cookie string) (*ory.Session, *http.Response, error) {
	return s.client.FrontendAPI.
		ToSession(ctx).
		Cookie(cookie).
		Execute()
}

// GetIdentity returns identity
func (s *KratosService) GetIdentity(ctx context.Context, id string) (*ory.Identity, *http.Response, error) {
	return s.client.IdentityAPI.GetIdentity(ctx, id).Execute()
}

// UpdateIdentityTraits updates identity traits
func (s *KratosService) UpdateIdentityTraits(ctx context.Context, id string, traits map[string]interface{}) (*ory.Identity, *http.Response, error) {
	payload := ory.UpdateIdentityBody{
		Traits: traits,
	}
	return s.client.IdentityAPI.UpdateIdentity(ctx, id).
		UpdateIdentityBody(payload).
		Execute()
}
