package services

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
	"gitlab.com/shaninalex/lumna/app2/pkg/token"
	"gitlab.com/shaninalex/lumna/app2/repositories"
)

type AuthManager interface {
	// Login generates new access + refresh tokens and stores refresh token
	// TODO: replace "device" argument with "meta" dictionary
	Login(ctx context.Context, userID uint) (*AccessTokenResult, *RefreshTokenResult, error)

	// Logout deletes refresh token (single device)
	Logout(ctx context.Context, userID uint, refreshToken string) error

	// ListSessions returns all refresh tokens for a user
	ListSessions(ctx context.Context, userID uint) ([]*models.UserToken, error)

	// RefreshAccessToken validates refresh token and returns a new access token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResult, error)
}

func NewAuthManager() AuthManager {
	return &AuthService{
		accessTokenService:  NewDefaultAccessTokenService(),
		refreshTokenService: NewDefaultRefreshTokenService(),
		userTokenRepository: repositories.NewUserTokenRepository(),
	}
}

type AuthService struct {
	accessTokenService  AccessTokenService
	refreshTokenService RefreshTokenService
	userTokenRepository *repositories.UserTokenRepository
}

var _ AuthManager = &AuthService{}

func (s *AuthService) Login(ctx context.Context, userId uint) (*AccessTokenResult, *RefreshTokenResult, error) {
	accessResult, err := s.accessTokenService.Create(userId, token.AudTokenAPIUser)
	if err != nil {
		return nil, nil, err
	}

	refreshResults, err := s.refreshTokenService.Create(userId)
	if err != nil {
		return nil, nil, err
	}

	tokenModel := &models.UserToken{
		UserId: userId,
		// Device:           device,
		RefreshToken:     refreshResults.Token,
		RefreshExpiresAt: refreshResults.ExpiresAt,
	}
	err = s.userTokenRepository.Create(ctx, tokenModel)
	if err != nil {
		return nil, nil, err
	}

	return accessResult, refreshResults, nil
}

func (s *AuthService) Logout(ctx context.Context, userID uint, refreshToken string) error {
	return s.userTokenRepository.DeleteTokenByRefreshString(ctx, userID, refreshToken)
}

func (s *AuthService) ListSessions(ctx context.Context, userID uint) ([]*models.UserToken, error) {
	return s.userTokenRepository.List(ctx, db.Eq("user_id", userID))
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResult, error) {
	if _, err := s.refreshTokenService.Validate(refreshToken); err != nil {
		return nil, err
	}

	t, err := s.userTokenRepository.GetTokenByField(ctx, "refresh_token", refreshToken)
	if err != nil {
		return nil, err
	}

	if t.IsRevoked() {
		return nil, fmt.Errorf("token is revoked")
	}

	if t.IsExpired() {
		if err := s.userTokenRepository.RevokeToken(ctx, t.UserId, t.GetId()); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("unable to revoke token")
		}
		return nil, fmt.Errorf("refresh token is expired")
	}

	result, err := s.accessTokenService.Create(t.UserId, token.AudTokenAPIUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
