package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_UserServiceGetUser(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewUserService()
	testUser := tests.CreateUser(ctx, "test@test.com")

	user, err := service.GetUser(ctx, testUser.GetId())
	assert.NoError(t, err, "should get user without errors")
	assert.Equal(t, user.Id, testUser.Id, "should have correct user id")

}

func Test_UserServiceGetUserByEmail(t *testing.T) {

}

func Test_UserServiceCheckPassword(t *testing.T) {

}

func Test_UserServiceCreateUser(t *testing.T) {

}

func Test_UserServiceUpdate(t *testing.T) {

}
