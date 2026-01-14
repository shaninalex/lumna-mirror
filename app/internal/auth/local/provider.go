package local

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

var _ auth.AuthProvider = (*LocalAuthProvider)(nil)

type PasswordCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *PasswordCredentials) Validate() error {
	return nil
}

type LocalAuthProvider struct {
}

func NewLocalAuthProvider() *LocalAuthProvider {
	s := &LocalAuthProvider{}
	return s
}

func (s *LocalAuthProvider) Name() string {
	return "local"
}

func (s *LocalAuthProvider) Authenticate(ctx context.Context, payload auth.AuthPayload) (*auth.AuthResult, error) {
	// cnf := config.GetConfig(ctx)
	// TODO:
	// find user credentials by Credential.Email
	// validate password
	// if error return nil, error
	// if not error:
	// 	- find Identity
	// 	- return Identity
	return nil, nil
}
