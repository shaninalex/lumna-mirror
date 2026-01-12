package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
)

type AuthProvider interface {
	Name() string
	Authenticate(ctx context.Context, payload any) (*models.Identity, error)
}
