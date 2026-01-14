package github

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

var _ auth.AuthProvider = (*GoogleAuthProvider)(nil)

type GoogleAuthProvider struct {
}

func (s *GoogleAuthProvider) Name() string {
	return "github"
}

func (s *GoogleAuthProvider) Authenticate(ctx context.Context, payload auth.AuthPayload) (*auth.AuthResult, error) {
	panic("not implenented")
}
