package token

import (
	"fmt"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/base"
)

// AccessTokenResult contains the signed token and metadata
type AccessTokenResult struct {
	Token     string
	ExpiresAt time.Time
	Sub       int64 // user Id
	JTI       string
}

// AccessTokenService handles creation and validation of access tokens
type AccessTokenService interface {
	// Create generates a new access token for a given user Id and audience
	Create(userID int64, aud AudToken) (*AccessTokenResult, error)

	// Validate parses and validates the token string, returns claims if valid
	Validate(rawToken string, aud AudToken) (*jwt.RegisteredClaims, error)
}

type AccessTokenJWTService struct {
	signingKey []byte
	issuer     string
}

func NewDefaultAccessTokenService() AccessTokenService {
	return NewAccessTokenJWTService(base.GetConfig().String("secret_key"), Issuer)
}

func NewAccessTokenJWTService(signingKey string, issuer string) *AccessTokenJWTService {
	return &AccessTokenJWTService{
		signingKey: []byte(signingKey),
		issuer:     issuer,
	}
}

func (s *AccessTokenJWTService) Create(userID int64, aud AudToken) (*AccessTokenResult, error) {
	now := time.Now()
	exp := now.Add(AccessTokenLifeTime)
	jti := uuid.NewString()
	claims := jwt.RegisteredClaims{
		Issuer:    s.issuer,
		Subject:   fmt.Sprintf("%d", userID),
		Audience:  []string{string(aud)},
		ExpiresAt: jwt.NewNumericDate(exp),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jti,
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.signingKey)
	if err != nil {
		return nil, err
	}

	return &AccessTokenResult{
		Token:     accessToken,
		ExpiresAt: exp,
		Sub:       userID,
		JTI:       jti,
	}, nil
}

func (s *AccessTokenJWTService) Validate(rawToken string, aud AudToken) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	if !slices.Contains(claims.Audience, string(aud)) {
		return nil, fmt.Errorf("invalid audience")
	}

	return claims, nil
}

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
	Create(userID int64, device string) (*RefreshTokenResult, error)

	// Validate checks if a given refresh token string is valid (signature/format)
	Validate(token string) (*jwt.RegisteredClaims, error)
}

type RefreshTokenJWTService struct {
	signingKey []byte
	issuer     string
}

func NewDefaultRefreshTokenService() RefreshTokenService {
	return NewRefreshTokenJWTService(base.GetConfig().String("secret_key"), Issuer)
}

func NewRefreshTokenJWTService(signingKey string, issuer string) *RefreshTokenJWTService {
	return &RefreshTokenJWTService{
		signingKey: []byte(signingKey),
		issuer:     issuer,
	}
}

func (s *RefreshTokenJWTService) Create(userID int64, device string) (*RefreshTokenResult, error) {
	now := time.Now()
	exp := now.Add(RefreshTokenLifeTime)
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
	token, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (interface{}, error) {
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
