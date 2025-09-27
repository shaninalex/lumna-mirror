package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/utils"
)

const (
	// AccessTokenLifeTime is a 15 minutes duration
	AccessTokenLifeTime = 15 * time.Minute

	// RefreshTokenLifeTime is a 7 days duration
	RefreshTokenLifeTime = 7 * 24 * time.Hour

	// RefreshTokenLength default refresh token string length
	RefreshTokenLength = 64

	// Issuer default issuer
	Issuer = "lumna"

	// NumericRefreshTokenLifeTime numeric refresh token life time
	NumericRefreshTokenLifeTime = 7 * 24 * 60 * 60

	// NumericAccessTokenLifeTime numeric access token life time
	NumericAccessTokenLifeTime = 15 * 60
)

type AudToken string

var (
	AudTokenAPIUser   AudToken = "api"
	AudTokenAdminUser AudToken = "admin"
	AudTokenExternal  AudToken = "external"
)

type TokenService interface {
	Create(userID uint, aud AudToken) (*TokenResult, error)
	Validate(rawToken string, aud AudToken) (jwt.Claims, error)
}

func NewDefaultTokenService() TokenService {
	key := utils.GetEnv("LUMNA_SECRET_KEY", "a-string-secret-at-least-256-bits-long")
	return NewJwtService(key, RefreshTokenLength, Issuer)
}

func NewJwtService(signingKey string, refreshStrLenght int, issuer string) *JwtService {
	return &JwtService{
		signingKey:       []byte(signingKey),
		refreshStrLenght: refreshStrLenght,
		issuer:           issuer,
	}
}

type TokenResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	RefreshExpAt time.Time
	Jti          string
	Sub          uint
}

type JwtService struct {
	signingKey       []byte
	refreshStrLenght int
	issuer           string
}

func (s *JwtService) Create(userID uint, aud AudToken) (*TokenResult, error) {
	now := time.Now()

	jti := uuid.NewString()
	exp := now.Add(AccessTokenLifeTime)
	claims := jwt.RegisteredClaims{
		Issuer:    s.issuer,
		Subject:   strconv.Itoa(int(userID)),
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

	refreshToken, err := generateSecureString(RefreshTokenLength)
	if err != nil {
		return nil, err
	}
	refreshExpiresAt := now.Add(RefreshTokenLifeTime)
	return &TokenResult{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		ExpiresAt:    exp,
		RefreshExpAt: refreshExpiresAt,
		Jti:          jti,
		Sub:          userID,
	}, nil
}

func (s *JwtService) Validate(rawToken string, aud AudToken) (jwt.Claims, error) {
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

func generateSecureString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
