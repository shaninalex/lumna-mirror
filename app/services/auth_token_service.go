package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
)

type AuthManager interface {
	// Login generates new access + refresh tokens and stores refresh token
	Login(ctx context.Context, userID uint, device string) (*AccessTokenResult, *RefreshTokenResult, error)

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
	}
}

type AuthService struct {
	accessTokenService  AccessTokenService
	refreshTokenService RefreshTokenService
}

func (s *AuthService) Login(ctx context.Context, userID uint, device string) (*AccessTokenResult, *RefreshTokenResult, error) {
	//accessResult, err := s.accessTokenService.Create(userID, token.AudTokenAPIUser)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	//refreshResults, err := s.refreshTokenService.Create(userID, device)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	//tokenModel := &models.UserToken{
	//	UserID:           userID,
	//	Device:           device,
	//	RefreshToken:     refreshResults.Token,
	//	RefreshExpiresAt: refreshResults.ExpiresAt,
	//}
	//err = domain.SaveToken(ctx, db.GetDb(ctx), tokenModel)
	//if err != nil {
	//	return nil, nil, err
	//}
	//return accessResult, refreshResults, nil
	panic("not implemented")
}

func (s *AuthService) Logout(ctx context.Context, userID uint, refreshToken string) error {
	//connection := db.GetDb(ctx)
	//return models.DeleteTokenByRefreshString(ctx, connection, userID, refreshToken)
	panic("not implemented")
}

func (s *AuthService) ListSessions(ctx context.Context, userID uint) ([]*models.UserToken, error) {
	//tokens, err := domain.GetTokens(ctx, db.GetDb(ctx), userID)
	//if err != nil {
	//	return nil, err
	//}
	//return tokens, nil
	panic("not implemented")
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (*AccessTokenResult, error) {
	//if _, err := s.refreshTokenService.Validate(refreshToken); err != nil {
	//	return nil, err
	//}
	//token, err := domain.GetTokenByField(ctx, db.GetDb(ctx), "refresh_token", refreshToken)
	//if err != nil {
	//	return nil, err
	//}
	//result, err := s.accessTokenService.Create(token.UserID, AudTokenAPIUser)
	//if err != nil {
	//	return nil, err
	//}
	//return result, nil
	panic("not implemented")
}
