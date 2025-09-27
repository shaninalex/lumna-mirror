package token_test

import (
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/internal/token"
)

var (
	signingKey       = "a-string-secret-at-least-256-bits-long"
	refreshStrLength = 64
	issuer           = "test_issuer"
	aud              = token.AudToken("test_audience")
)

func Test_CreateToken(t *testing.T) {
	service := token.NewJwtService(signingKey, refreshStrLength, issuer)
	userId := uint(100)
	result, err := service.Create(userId, aud)

	assert.Nil(t, err)
	assert.NotEqual(t, "", result.AccessToken)
	assert.NotEqual(t, "", result.RefreshToken)
	assert.NotEqual(t, "", result.Jti)

	assert.Equal(t, userId, result.Sub)
	assert.NotEqual(t, result.AccessToken, result.RefreshToken)

	jti, err := uuid.Parse(result.Jti)
	assert.Nil(t, err)
	assert.NotEqual(t, uuid.Nil, jti)
}

func Test_ValidateToken(t *testing.T) {
	service := token.NewJwtService(signingKey, refreshStrLength, issuer)
	userId := uint(100)
	result, _ := service.Create(userId, aud)
	claims, err := service.Validate(result.AccessToken, aud)

	assert.Nil(t, err)

	registeredClaims, ok := claims.(*jwt.RegisteredClaims)

	assert.True(t, ok)
	assert.Equal(t, strconv.Itoa(int(userId)), registeredClaims.Subject)
	assert.Equal(t, issuer, registeredClaims.Issuer)
	assert.True(t, slices.Contains(registeredClaims.Audience, string(aud)))

	now := time.Now()
	assert.NotNil(t, registeredClaims.IssuedAt)
	assert.WithinDuration(t, now, registeredClaims.IssuedAt.Time, 2*time.Second)

	assert.NotNil(t, registeredClaims.NotBefore)
	assert.True(t, !registeredClaims.NotBefore.Time.After(now))

	assert.NotNil(t, registeredClaims.ExpiresAt)
	assert.True(t, registeredClaims.ExpiresAt.Time.After(now))
	assert.WithinDuration(t, now.Add(token.AccessTokenLifeTime), registeredClaims.ExpiresAt.Time, 2*time.Second)
}
