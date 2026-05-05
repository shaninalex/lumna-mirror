package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/auth"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_AuthTokenService_Create(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	user := testutils.User(models.Identity{FullName: "test", Email: "test@test.com", Active: true}, db)
	ts := services.NewAuthTokenService(
		repositories.NewGormRefreshTokenRepository(db),
	)

	_, toDB, err := auth.GenerateRefreshToken()
	assert.NoError(t, err)

	refreshTtl := time.Hour * 24 * 30 // 30 days
	refreshExp := time.Now().Add(refreshTtl)
	rt := models.RefreshToken{
		IdentityID: user.ID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	err = ts.CreateRefreshToken(ctx, user.ID, &rt)
	assert.NoError(t, err)

	token, err := ts.GetByHash(ctx, rt.Hash)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	assert.Equal(t, rt.IdentityID, token.IdentityID)
	assert.Equal(t, rt.Hash, token.Hash)
	assert.Equal(t, rt.ClientID, token.ClientID)
	assert.Equal(t, rt.Scopes, token.Scopes)
	assert.Equal(t, rt.ExpiresAt, refreshExp)
}

func Test_AuthTokenService_Delete(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	user := testutils.User(models.Identity{FullName: "test", Email: "test@test.com", Active: true}, db)

	ts := services.NewAuthTokenService(
		repositories.NewGormRefreshTokenRepository(db),
	)
	_, toDB, _ := auth.GenerateRefreshToken()
	refreshTtl := time.Hour * 24 * 30 // 30 days
	refreshExp := time.Now().Add(refreshTtl)
	rt := models.RefreshToken{
		IdentityID: user.ID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	_ = ts.CreateRefreshToken(ctx, user.ID, &rt)

	err := ts.DeleteRefreshToken(ctx, user.ID)
	assert.NoError(t, err)

	token, err := ts.GetByHash(ctx, rt.Hash)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "record not found")
	assert.Nil(t, token)
}
