package oauth

import (
	"context"
	"time"
)

type RefreshToken struct {
	Token     string
	ClientID  string
	UserID    string
	Scopes    []string
	ExpiresAt time.Time
	Revoked   bool
}

type RefreshTokenStore interface {
	Save(ctx context.Context, token *RefreshToken) error
}
