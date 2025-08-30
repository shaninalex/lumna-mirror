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
	GetLoginFlow(ctx context.Context, cookie, flowID string) (*ory.LoginFlow, *http.Response, error)
	GetRegistrationFlow(ctx context.Context, cookie, flowID string) (*ory.RegistrationFlow, *http.Response, error)
	GetErrorFlow(ctx context.Context, errorID string) (*ory.FlowError, *http.Response, error)
	GetVerificationFlow(ctx context.Context, cookie, flowID string) (*ory.VerificationFlow, *http.Response, error)
	CreateLogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, *http.Response, error)
	CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, *http.Response, error)
	GetSettingsFlow(ctx context.Context, cookie, flowID string) (*ory.SettingsFlow, *http.Response, error)
	GetRecoveryFlow(ctx context.Context, cookie, flowID string) (*ory.RecoveryFlow, *http.Response, error)
	GetSession(ctx context.Context, cookie string) (*ory.Session, *http.Response, error)
	GetIdentity(ctx context.Context, id string) (*ory.Identity, *http.Response, error)
	UpdateIdentityTraits(ctx context.Context, id string, traits map[string]interface{}) (*ory.Identity, *http.Response, error)
}

// Service - kratos service.
type Service struct {
	client *ory.APIClient
}

// NewKratosService - new kratos service.
func NewKratosService(url string) *Service {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: url,
		},
	}
	return &Service{
		client: ory.NewAPIClient(configuration),
	}
}

// Client return ory kratos api client itself
func (s *Service) Client() *ory.APIClient {
	return s.client
}

// GetLoginFlow is the method that return created login flow based on user
// cookies and flow id
func (s *Service) GetLoginFlow(ctx context.Context, cookie, flowID string) (*ory.LoginFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetLoginFlow(ctx).
		Cookie(cookie).
		Id(flowID).
		Execute()
}

// GetRegistrationFlow is the method that return registration flow based on user
// cookies and flow id
func (s *Service) GetRegistrationFlow(ctx context.Context, cookie, flowID string) (*ory.RegistrationFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetRegistrationFlow(ctx).
		Cookie(cookie).
		Id(flowID).
		Execute()
}

// GetErrorFlow is the method that return error flow based on user
// cookies and flow id. ory/kratos store errors in database and user can
// access them via error flow id.
func (s *Service) GetErrorFlow(ctx context.Context, errorID string) (*ory.FlowError, *http.Response, error) {
	return s.client.FrontendAPI.
		GetFlowError(ctx).
		Id(errorID).
		Execute()
}

// GetVerificationFlow is the method that return verification flow based on user
// cookies and flow id. After registration user can be redirected to the
// verification page where this form should be rendered.
func (s *Service) GetVerificationFlow(ctx context.Context, cookie, flowID string) (*ory.VerificationFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetVerificationFlow(ctx).
		Cookie(cookie).
		Id(flowID).
		Execute()
}

// CreateLogoutFlow is the method that create browser logout flow. This flow
// return logout token and logout url. By following logout id or calling api
// call with logout token - kratos will remove active session from db and user
// will be logged out
func (s *Service) CreateLogoutFlow(ctx context.Context, cookie string) (*ory.LogoutFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		CreateBrowserLogoutFlow(ctx).
		Cookie(cookie).
		Execute()
}

// CreateSettingsFlow is the method that create user settings flow. This method
// basically says to the system that this user ( by cookie ) want to get his
// settings form to change something about himself.
func (s *Service) CreateSettingsFlow(ctx context.Context, cookie string) (*ory.SettingsFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		CreateBrowserSettingsFlow(ctx).
		Cookie(cookie).
		Execute()
}

// GetSettingsFlow is the method that return settings flow to user. This object
// contain all editable user profile fields that kratos user have access to.
func (s *Service) GetSettingsFlow(ctx context.Context, cookie, flowID string) (*ory.SettingsFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetSettingsFlow(ctx).
		Cookie(cookie).
		Id(flowID).
		Execute()
}

// GetRecoveryFlow is the method that return account recovery flow
func (s *Service) GetRecoveryFlow(ctx context.Context, cookie, flowID string) (*ory.RecoveryFlow, *http.Response, error) {
	return s.client.FrontendAPI.
		GetRecoveryFlow(ctx).
		Cookie(cookie).
		Id(flowID).
		Execute()
}

// GetSession is the method that return user session information
func (s *Service) GetSession(ctx context.Context, cookie string) (*ory.Session, *http.Response, error) {
	return s.client.FrontendAPI.
		ToSession(ctx).
		Cookie(cookie).
		Execute()
}

// GetIdentity returns identity
func (s *Service) GetIdentity(ctx context.Context, id string) (*ory.Identity, *http.Response, error) {
	return s.client.IdentityAPI.GetIdentity(ctx, id).Execute()
}

// UpdateIdentityTraits updates identity traits
func (s *Service) UpdateIdentityTraits(ctx context.Context, id string, traits map[string]interface{}) (*ory.Identity, *http.Response, error) {
	payload := ory.UpdateIdentityBody{
		Traits: traits,
	}
	return s.client.IdentityAPI.UpdateIdentity(ctx, id).
		UpdateIdentityBody(payload).
		Execute()
}
