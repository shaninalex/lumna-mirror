// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_UserService_GetUser(t *testing.T) {
	m := tdata.Manager()
	testUser := tdata.CreateUser(m.Ctx)
	service := domain.NewUserService()
	user, err := service.GetUser(m.Ctx, testUser.GetID())
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, testUser.GetCode(), user.GetCode())
	assert.Equal(t, testUser.GetIdentity().GetId(), testUser.GetIdentity().GetId())
}

func Test_UserService_UpdateUserSettings(t *testing.T) {
	m := tdata.Manager()
	testUser := tdata.CreateUser(m.Ctx)
	service := domain.NewUserService()

	settings := testUser.GetSettings()
	settings.Language = "fr"

	err := service.UpdateUserSettings(m.Ctx, testUser.GetID(), settings)
	assert.NoError(t, err)

	var user models.User
	database.GetDB(m.Ctx).First(&user, testUser.GetID())
	assert.Equal(t, settings, user.GetSettings())
	assert.Equal(t, "fr", user.GetSettings().Language)
}
