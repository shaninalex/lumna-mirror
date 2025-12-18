package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_ServiceUser_Get(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewUserManager()
	testUser := tests.CreateUser(ctx, "test@test.com")

	user, err := service.GetUser(ctx, testUser.GetId())
	assert.NoError(t, err, "should get user without errors")
	assert.Equal(t, user.Id, testUser.Id, "should have correct user id")
}

func Test_ServiceUser_GetByEmail(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewUserManager()
	testUser := tests.CreateUser(ctx, "test@test.com")

	user, err := service.GetUserByEmail(ctx, testUser.Email)
	assert.NoError(t, err, "should get user without errors")
	assert.Equal(t, user.Email, testUser.Email, "should have correct user email")
}

func Test_ServiceUser_CheckPassword(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewUserManager()
	testUser := tests.CreateUser(ctx, "test@test.com")

	err := service.CheckPassword(ctx, testUser.GetId(), testUser.Email)
	assert.NoError(t, err, "should check password without errors")
}

func Test_ServiceUser_CreateUser(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	service := services.NewUserManager()

	user, err := service.CreateUser(ctx, "test@test.com", "password")
	assert.NoError(t, err, "should create user without errors")
	err = service.CheckPassword(ctx, user.GetId(), "password")
	assert.NoError(t, err, "should check password without errors")
}

func Test_ServiceUser_Update(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewUserManager()
	testUser := tests.CreateUser(ctx, "test@test.com")

	newEmail := "test2@test.com"
	err := service.Update(ctx, testUser.GetId(), db.Option{Key: "email", Value: newEmail})
	assert.NoError(t, err, "should update user without errors")

	user, err := service.GetUser(ctx, testUser.GetId())
	assert.NoError(t, err, "should get user without errors")
	assert.Equal(t, user.Email, newEmail, "should have correct user email")
}
