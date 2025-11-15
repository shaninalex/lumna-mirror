package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type UserToken struct {
	Id               int64      `db:"id"`
	UserID           int64      `db:"user_id"`
	Device           string     `db:"device"`
	RefreshToken     string     `db:"refresh_token"`
	RefreshExpiresAt time.Time  `db:"refresh_expires_at"`
	Revoked          bool       `db:"revoked"`
	RevokedAt        *time.Time `db:"revoked_at"`
	CreatedAt        time.Time  `db:"created_at"`
}

// GetID - returns the id.
func (s *UserToken) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *UserToken) SetID(id int64) { s.Id = id }

// IsExpired - check is token expired
func (s *UserToken) IsExpired() bool {
	return s.RefreshExpiresAt.Before(time.Now())
}

// UserTokenManager defines the interface for managing user tokens.
// It abstracts operations like listing, deleting, or revoking tokens.
// Implementations should not care about how tokens are stored.
type UserTokenManager interface {
	// List returns all tokens associated with a specific user.
	List(ctx context.Context, userID int64) ([]*UserToken, error)

	// Delete removes a token record permanently from the database.
	Delete(ctx context.Context, userID, tokenID int64) error

	// Revoke marks a token as revoked without deleting it.
	// Useful for invalidating a refresh token or access session.
	Revoke(ctx context.Context, userID, tokenID int64) error
}

// UserTokenService is a concrete implementation of UserTokenManager.
// It uses the repositories layer to interact with the database.
type UserTokenService struct{}

// NewUserTokenService creates a new instance of UserTokenService.
func NewUserTokenService() *UserTokenService {
	return &UserTokenService{}
}

// List fetches all tokens for a given user from the database.
// It calls the repositories layer and returns any database errors.
func (u UserTokenService) List(ctx context.Context, userID int64) ([]*UserToken, error) {
	tokens, err := GetTokens(ctx, db.GetDb(ctx), userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// Delete removes a token for a specific user from the database.
// Typically used for logout or administrative token cleanup.
func (u UserTokenService) Delete(ctx context.Context, userID, tokenID int64) error {
	return DeleteToken(ctx, db.GetDb(ctx), userID, tokenID)
}

// Revoke invalidates a token without deleting it from the database.
// Useful for forcing logouts or invalidating refresh tokens.
func (u UserTokenService) Revoke(ctx context.Context, userID, tokenID int64) error {
	return RevokeToken(ctx, db.GetDb(ctx), userID, tokenID)
}
