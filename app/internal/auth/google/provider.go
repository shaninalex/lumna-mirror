package github

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/models"
)

var _ auth.AuthProvider = (*GoogleAuthProvider)(nil)

type GoogleAuthProvider struct {
}

func (s *GoogleAuthProvider) Name() string {
	return "github"
}

func (s *GoogleAuthProvider) Authenticate(ctx context.Context, payload any) (*models.Identity, error) {
	panic("not implenented")
}
