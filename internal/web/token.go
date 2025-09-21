// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

type TokenManager interface {
	CreateToken(ctx context.Context, userID uuid.UUID) (string, error)
	ValidateToken(ctx context.Context, rawToken string) (jwt.MapClaims, error)
	GetTokens(ctx context.Context, userID uuid.UUID) ([]*models.UserToken, error)
	DeleteToken(ctx context.Context, userID, tokenID uuid.UUID) error
}

var sampleSecretKey = []byte("GoLinuxCloudKey")

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

// CreateToken - create token
func (s *TokenService) CreateToken(ctx context.Context, userID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	now := time.Now()
	exp := now.Add(time.Minute * 30)
	jti := uuid.New()
	claims["jti"] = jti.String()
	claims["iat"] = now.Unix()
	claims["sub"] = userID.String()
	claims["exp"] = exp.Unix()
	claims["roles"] = []string{"user"}
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	claimsMap := map[string]any{}
	for k, v := range claims {
		claimsMap[k] = v
	}
	device := ctx.Value("device").(string)
	tokenModel := &models.UserToken{
		ID:        jti,
		UserID:    userID,
		Device:    device,
		ExpiresAt: exp,
		CreatedAt: now,
	}

	tokenModel.SetClaims(claimsMap)
	err = repositories.SaveToken(ctx, database.GetDb(ctx), tokenModel)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken - validate given raw access token
func (s *TokenService) ValidateToken(ctx context.Context, rawToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, fmt.Errorf("there was an error in parsing")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	if err = s.claimsValidation(ctx, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// GetTokens - get tokens from database
func (s *TokenService) GetTokens(ctx context.Context, userID uuid.UUID) ([]*models.UserToken, error) {
	return repositories.GetTokens(ctx, database.GetDb(ctx), userID)
}

// DeleteToken - delete token from database
func (s *TokenService) DeleteToken(ctx context.Context, userID, tokenID uuid.UUID) error {
	return repositories.DeleteToken(ctx, database.GetDb(ctx), userID, tokenID)
}

// claimsValidation - token claims validation
func (s *TokenService) claimsValidation(ctx context.Context, claims jwt.MapClaims) error {
	// Check jti exists in DB
	jti, ok := claims["jti"].(string)
	if !ok || jti == "" {
		return fmt.Errorf("token missing jti claim")
	}

	userIDStr, ok := claims["sub"].(string)
	if !ok || userIDStr == "" {
		return fmt.Errorf("token missing sub claim")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("invalid user_id in token")
	}

	tokenID, err := uuid.Parse(jti)
	if err != nil {
		return fmt.Errorf("invalid token id")
	}

	// Example DB lookup (replace with your actual DB call)
	dbToken, err := s.getTokenFromDB(ctx, userID, tokenID)
	if err != nil {
		return fmt.Errorf("token not found or revoked")
	}
	if dbToken.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("token expired in DB")
	}
	return nil
}

// getTokenFromDB get token from database1
func (s *TokenService) getTokenFromDB(ctx context.Context, userID, tokenID uuid.UUID) (*models.UserToken, error) {
	db := database.GetDb(ctx)
	token, err := repositories.GetTokenByField(ctx, db, "id", tokenID.String())
	if err != nil {
		return nil, err
	}
	if token.UserID != userID {
		return nil, fmt.Errorf("invalid user id")
	}
	return token, nil
}
