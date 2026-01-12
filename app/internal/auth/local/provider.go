package local

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/models"
)

var _ auth.AuthProvider = (*LocalAuthProvider)(nil)

type LocalAuthProvider struct {
}

func (s *LocalAuthProvider) Name() string {
	return "local"
}

func (s *LocalAuthProvider) Authenticate(ctx context.Context, payload any) (*models.Identity, error) {
	return nil, nil
}
