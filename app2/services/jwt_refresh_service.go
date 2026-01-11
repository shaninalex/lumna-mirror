package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app2/pkg/token"
)

// RefreshTokenResult contains the generated token and its expiration
type RefreshTokenResult struct {
	Token     string
	ExpiresAt time.Time
}

type RefreshTokenClaims struct {
	Device string `json:"device"`
	jwt.RegisteredClaims
}

// RefreshTokenService handles creation and validation of refresh tokens
type RefreshTokenService interface {
	// Create generates a new refresh token string with expiration
	Create(userID uint) (*RefreshTokenResult, error)

	// Validate checks if a given refresh token string is valid (signature/format)
	Validate(token string) (*jwt.RegisteredClaims, error)

	// RefreshToken
}

type RefreshTokenJWTService struct {
	signingKey []byte
	issuer     string
}

var _ RefreshTokenService = &RefreshTokenJWTService{}

func NewDefaultRefreshTokenService() RefreshTokenService {
	return NewRefreshTokenJWTService(secretKey, token.Issuer)
}

func NewRefreshTokenJWTService(signingKey string, issuer string) *RefreshTokenJWTService {
	return &RefreshTokenJWTService{
		signingKey: []byte(signingKey),
		issuer:     issuer,
	}
}

func (s *RefreshTokenJWTService) Create(userID uint) (*RefreshTokenResult, error) {
	now := time.Now()
	exp := now.Add(token.RefreshTokenLifeTime)
	jti := uuid.NewString()
	claims := jwt.RegisteredClaims{
		Issuer:    s.issuer,
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jti,
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.signingKey)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenResult{
		Token:     refreshToken,
		ExpiresAt: exp,
	}, nil
}

func (s *RefreshTokenJWTService) Validate(rawToken string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check token validity
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check expiration
	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
		return nil, fmt.Errorf("token expired")
	}

	// Validate jti is a UUID
	if _, err := uuid.Parse(claims.ID); err != nil {
		return nil, fmt.Errorf("invalid jti: %w", err)
	}

	return claims, nil
}
