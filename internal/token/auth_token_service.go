// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package token

import (
	"context"

	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

type ApiAuthService interface {
	// Login generates new access + refresh tokens and stores refresh token
	Login(ctx context.Context, userID uint, device string) (*AccessTokenResult, *RefreshTokenResult, error)

	// Logout deletes refresh token (single device)
	Logout(ctx context.Context, userID uint, refreshToken string) error

	// ListSessions returns all refresh tokens for a user
	ListSessions(ctx context.Context, userID uint) ([]*models.UserToken, error)

	// RefreshAccessToken validates refresh token and returns a new access token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResult, error)
}

func NewAuthService() ApiAuthService {
	return &AuthService{
		accessTokenService:  NewDefaultAccessTokenService(),
		refreshTokenService: NewDefaultRefreshTokenService(),
	}
}

type AuthService struct {
	accessTokenService  AccessTokenService
	refreshTokenService RefreshTokenService
}

func (s *AuthService) Login(ctx context.Context, userID uint, device string) (*AccessTokenResult, *RefreshTokenResult, error) {
	accessResult, err := s.accessTokenService.Create(userID, AudTokenAPIUser)
	if err != nil {
		return nil, nil, err
	}

	refreshResults, err := s.refreshTokenService.Create(userID, device)
	if err != nil {
		return nil, nil, err
	}

	tokenModel := &models.UserToken{
		UserID:           userID,
		Device:           device,
		RefreshToken:     refreshResults.Token,
		RefreshExpiresAt: refreshResults.ExpiresAt,
	}
	db := database.GetDb(ctx)
	err = repositories.SaveToken(ctx, db, tokenModel)
	if err != nil {
		return nil, nil, err
	}
	return accessResult, refreshResults, nil
}

func (s *AuthService) Logout(ctx context.Context, userID uint, refreshToken string) error {
	db := database.GetDb(ctx)
	return repositories.DeleteTokenByRefreshString(ctx, db, userID, refreshToken)
}

func (s *AuthService) ListSessions(ctx context.Context, userID uint) ([]*models.UserToken, error) {
	db := database.GetDb(ctx)
	tokens, err := repositories.GetTokens(ctx, db, userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResult, error) {
	db := database.GetDb(ctx)
	if _, err := s.refreshTokenService.Validate(refreshToken); err != nil {
		return nil, err
	}
	token, err := repositories.GetTokenByField(ctx, db, "refresh_token", refreshToken)
	if err != nil {
		return nil, err
	}
	result, err := s.accessTokenService.Create(token.UserID, AudTokenAPIUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
