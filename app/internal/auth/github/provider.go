package github

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

var _ auth.AuthProvider = (*GithubAuthProvider)(nil)

type GithubAuthProvider struct {
}

func (s *GithubAuthProvider) Name() string {
	return "github"
}

func (s *GithubAuthProvider) Authenticate(ctx context.Context, payload auth.AuthPayload) (*auth.AuthResult, error) {
	panic("not implenented")
}
