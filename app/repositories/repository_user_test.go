package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_CreateUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()

	user := models.User{}

	err := repo.Create(ctx, &user)

	assert.NoError(t, err, "Should create user without errors")
	assert.True(t, user.Id > 0, "User should have not empty id")
}
