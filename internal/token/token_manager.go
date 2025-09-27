// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package token

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

type TokenManager interface {
	Create(ctx context.Context, userID uint) (string, string, error)
	Validate(ctx context.Context, rawToken string) (*jwt.RegisteredClaims, error)
	List(ctx context.Context, userID uint) ([]*models.UserToken, error)
	Delete(ctx context.Context, userID uint, tokenID string) error
}

var sampleSecretKey = []byte(utils.GetEnv("FLOWREON_SECRET_KEY", "a-string-secret-at-least-256-bits-long"))
var expDelta = 7 * 24 * time.Hour // 1 week

type tokenManager struct {
	tokenService TokenService
}

func NewTokenManager() TokenManager {
	return &tokenManager{
		tokenService: NewDefaultTokenService(),
	}
}

// Create - create token
func (s *tokenManager) Create(ctx context.Context, userID uint) (string, string, error) {
	result, err := s.tokenService.Create(userID, AudTokenAPIUser)
	if err != nil {
		return "", "", err
	}
	device := ctx.Value("device").(string)
	tokenModel := &models.UserToken{
		UserID:           userID,
		Device:           device,
		Jti:              result.Jti,
		RefreshToken:     result.RefreshToken,
		RefreshExpiresAt: result.RefreshExpAt,
	}
	err = repositories.SaveToken(ctx, database.GetDb(ctx), tokenModel)
	if err != nil {
		return "", "", err
	}

	return result.AccessToken, result.RefreshToken, nil
}

// Validate - validate given raw access token
func (s *tokenManager) Validate(ctx context.Context, rawToken string) (*jwt.RegisteredClaims, error) {
	clms, err := s.tokenService.Validate(rawToken, AudTokenAPIUser)
	if err != nil {
		return nil, err
	}

	claims, ok := clms.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	if err = s.claimsValidation(ctx, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// List - get tokens from database
func (s *tokenManager) List(ctx context.Context, userID uint) ([]*models.UserToken, error) {
	return repositories.GetTokens(ctx, database.GetDb(ctx), userID)
}

// Delete - delete token from database
func (s *tokenManager) Delete(ctx context.Context, userID uint, tokenID string) error {
	return repositories.DeleteToken(ctx, database.GetDb(ctx), userID, tokenID)
}

// claimsValidation - token claims validation
func (s *tokenManager) claimsValidation(ctx context.Context, claims *jwt.RegisteredClaims) error {
	// Check jti exists in DB
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	// Example DB lookup (replace with your actual DB call)
	_, err = s.getTokenFromDB(ctx, uint(userID), claims.ID)
	if err != nil {
		return fmt.Errorf("token not found or revoked")
	}
	return nil
}

// getTokenFromDB get token from database1
func (s *tokenManager) getTokenFromDB(ctx context.Context, userID uint, jti string) (*models.UserToken, error) {
	db := database.GetDb(ctx)
	token, err := repositories.GetTokenByField(ctx, db, "jti", jti)
	if err != nil {
		return nil, err
	}
	if token.UserID != userID {
		return nil, fmt.Errorf("invalid user id")
	}
	return token, nil
}

const (
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"
)

// ClearAuthCookies clear all auth cookies from response to unauthenticate user
func ClearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // expire immediately
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // expire immediately
		SameSite: http.SameSiteStrictMode,
	})
}
