package services_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/services"
)

func Test_ServiceAccessToken_CreateAndValidate(t *testing.T) {
	service := services.NewDefaultAccessTokenService()
	var userId uint = 1
	accessToken, err := service.Create(userId, token.AudTokenAPIUser)
	assert.NoError(t, err)
	assert.Equal(t, userId, accessToken.Sub)
	claims, err := service.Validate(accessToken.Token, token.AudTokenAPIUser)
	assert.NoError(t, err)
	id, err := strconv.Atoi(claims.Subject)
	assert.NoError(t, err)
	assert.Equal(t, userId, uint(id))

	b, err := claims.Audience.MarshalJSON()
	assert.NoError(t, err)
	var auds []string
	err = json.Unmarshal(b, &auds)
	assert.Equal(t, 1, len(auds))
	assert.NoError(t, err)
	assert.Equal(t, string(token.AudTokenAPIUser), auds[0])
}
