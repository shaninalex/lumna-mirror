package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
)

type AuthResult struct {
	Identity   *models.Identity
	Provider   string
	Credential *models.Credential
}

type AuthProvider interface {
	Name() string
	Authenticate(ctx context.Context, payload AuthPayload) (*AuthResult, error)
}

type AuthPayload interface {
	Validate() error
}
