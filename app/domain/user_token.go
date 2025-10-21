// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"

	"github.com/shaninalex/lumna/app/internal/db"
)

// UserTokenManager defines the interface for managing user tokens.
// It abstracts operations like listing, deleting, or revoking tokens.
// Implementations should not care about how tokens are stored.
type UserTokenManager interface {
	// List returns all tokens associated with a specific user.
	List(ctx context.Context, userID uint) ([]*db.UserToken, error)

	// Delete removes a token record permanently from the database.
	Delete(ctx context.Context, userID, tokenID uint) error

	// Revoke marks a token as revoked without deleting it.
	// Useful for invalidating a refresh token or access session.
	Revoke(ctx context.Context, userID, tokenID uint) error
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
func (u UserTokenService) List(ctx context.Context, userID uint) ([]*db.UserToken, error) {
	tokens, err := db.GetTokens(ctx, db.GetDb(ctx), userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// Delete removes a token for a specific user from the database.
// Typically used for logout or administrative token cleanup.
func (u UserTokenService) Delete(ctx context.Context, userID, tokenID uint) error {
	return db.DeleteToken(ctx, db.GetDb(ctx), userID, tokenID)
}

// Revoke invalidates a token without deleting it from the database.
// Useful for forcing logouts or invalidating refresh tokens.
func (u UserTokenService) Revoke(ctx context.Context, userID, tokenID uint) error {
	return db.RevokeToken(ctx, db.GetDb(ctx), userID, tokenID)
}
