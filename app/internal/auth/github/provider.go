package github

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/models"
)

var _ auth.AuthProvider = (*GithubAuthProvider)(nil)

type GithubAuthProvider struct {
}

func (s *GithubAuthProvider) Name() string {
	return "github"
}

func (s *GithubAuthProvider) Authenticate(ctx context.Context, payload any) (*models.Identity, error) {
	panic("not implenented")
}
