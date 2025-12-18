package services_test

import (
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_ServiceAuth_Login(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	testUser := tests.CreateUser(ctx, "test@test.com")
	authManager := services.NewAuthManager()
	access, refresh, err := authManager.Login(ctx, testUser.GetId())
	assert.NoError(t, err)
	assert.NotNil(t, access)
	assert.NotNil(t, refresh)

	assert.Equal(t, access.Sub, testUser.GetId())
	assert.True(t, access.ExpiresAt.After(time.Now()))

	repo := repositories.NewUserTokenRepository()

	dbToken, err := repo.GetTokenByField(ctx, "user_id", testUser.GetId())
	assert.NoError(t, err)
	assert.NotNil(t, dbToken)
	assert.False(t, dbToken.IsRevoked())
	assert.False(t, dbToken.IsExpired())

	accessService := services.NewDefaultAccessTokenService()
	claims, err := accessService.Validate(access.Token, token.AudTokenAPIUser)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	uintSubject, err := strconv.Atoi(claims.Subject)
	assert.NoError(t, err)
	assert.Equal(t, uint(uintSubject), testUser.GetId())

	refreshService := services.NewDefaultRefreshTokenService()
	claims, err = refreshService.Validate(refresh.Token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	uintSubject, err = strconv.Atoi(claims.Subject)
	assert.NoError(t, err)
	assert.Equal(t, uint(uintSubject), testUser.GetId())
}

func Test_ServiceAuth_Refresh(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	testUser := tests.CreateUser(ctx, "test@test.com")
	authManager := services.NewAuthManager()
	_, refresh, _ := authManager.Login(ctx, testUser.GetId())
	access, err := authManager.RefreshAccessToken(ctx, refresh.Token)
	assert.NoError(t, err)
	assert.NotNil(t, access)
}

func Test_ServiceAuth_Logout(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	testUser := tests.CreateUser(ctx, "test@test.com")
	authManager := services.NewAuthManager()
	_, refresh, _ := authManager.Login(ctx, testUser.GetId())
	err := authManager.Logout(ctx, testUser.GetId(), refresh.Token)
	assert.NoError(t, err)

	repo := repositories.NewUserTokenRepository()
	_, err = repo.GetTokenByField(ctx, "user_id", testUser.GetId())
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

func Test_ServiceAuth_ListSessions(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	testUser := tests.CreateUser(ctx, "test@test.com")
	authManager := services.NewAuthManager()
	_, _, _ = authManager.Login(ctx, testUser.GetId())

	list, err := authManager.ListSessions(ctx, testUser.GetId())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))
	assert.Equal(t, list[0].UserId, testUser.GetId())
}
