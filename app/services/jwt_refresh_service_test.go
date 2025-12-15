package services_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/services"
)

func Test_RefreshTokenServiceCreateAndValidate(t *testing.T) {
	service := services.NewDefaultRefreshTokenService()
	var userId uint = 1
	refreshToken, err := service.Create(userId)
	assert.NoError(t, err)
	claims, err := service.Validate(refreshToken.Token)
	assert.NoError(t, err)
	assert.Equal(t, claims.Subject, claims.Subject)
	id, err := strconv.Atoi(claims.Subject)
	assert.NoError(t, err)
	assert.Equal(t, userId, uint(id))
}
