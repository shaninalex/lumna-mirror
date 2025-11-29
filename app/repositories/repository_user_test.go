package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_CreateUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()

	h, _ := utils.CreatePasswordHash("123")
	user := models.User{
		Email:        "test@test.com",
		PasswordHash: h,
	}

	err := repo.Create(ctx, &user)
	assert.NoError(t, err, "Should create user without errors")
	if err != nil {
		t.Log(err)
		return
	}
	assert.True(t, user.Id > 0, "User should have not empty id")
	assert.True(t, !user.CreatedAt.IsZero(), "User should not have zero creation time")
}

func Test_UpdateUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	h, _ := utils.CreatePasswordHash("123")
	user := models.User{
		Email:        "test@test.com",
		PasswordHash: h,
	}

	newEmail := "test2@test.com"

	err := repo.Create(ctx, &user)
	if err != nil {
		t.Log(err)
		return
	}
	err = repo.Update(ctx, &user, map[string]any{
		"email": newEmail,
	})
	assert.NoError(t, err, "Should update user without errors")
	if err != nil {
		t.Log(err)
		return
	}

	assert.True(t, user.Email == newEmail, "User should have updated email")
}
