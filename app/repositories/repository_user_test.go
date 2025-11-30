package repositories_test

import (
	"context"
	"fmt"
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
	err = repo.Update(ctx, &user, repositories.Option{Key: "email", Value: newEmail})
	assert.NoError(t, err, "Should update user without errors")
	if err != nil {
		t.Log(err)
		return
	}

	assert.True(t, user.Email == newEmail, "User should have updated email")
}

func Test_GetUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	h, _ := utils.CreatePasswordHash("123")
	user := models.User{
		Email:        "test@test.com",
		PasswordHash: h,
	}
	_ = repo.Create(ctx, &user)

	dbUser, err := repo.Get(ctx, user.Id)
	assert.NoError(t, err, "should get user by id")
	assert.Equal(t, dbUser.Id, user.Id)
	assert.Equal(t, dbUser.Email, user.Email)
}

func Test_GetUserByEmail(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	h, _ := utils.CreatePasswordHash("123")
	user := models.User{
		Email:        "test@test.com",
		PasswordHash: h,
	}
	_ = repo.Create(ctx, &user)

	dbUser, err := repo.GetByEmail(ctx, user.Email)
	assert.NoError(t, err, "should get user by email")
	assert.Equal(t, dbUser.Id, user.Id)
	assert.Equal(t, dbUser.Email, user.Email)
}

func Test_DeleteUser(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	h, _ := utils.CreatePasswordHash("123")
	user := models.User{
		Email:        "test@test.com",
		PasswordHash: h,
	}
	_ = repo.Create(ctx, &user)

	err := repo.Delete(ctx, user.Id)
	assert.NoError(t, err, "should get user by email")

	_, err = repo.Get(ctx, user.Id)
	assert.Error(t, err, "should not get user")
}

func Test_ListUsers(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	userA := createUser(ctx, repo, "test1@test.com")
	_ = createUser(ctx, repo, "test2@test.com")

	usersList, err := repo.List(ctx, fmt.Sprintf("email = \"%s\"", userA.Email))
	assert.NoError(t, err, "should list users without errors")
	assert.Equal(t, len(usersList), 1, "should get only 1 user")
	assert.Equal(t, usersList[0].Email, userA.Email)
}

func Test_ListAllUsers(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	userA := createUser(ctx, repo, "test1@test.com")
	userB := createUser(ctx, repo, "test2@test.com")

	usersList, err := repo.List(ctx, "")
	assert.NoError(t, err, "should list users without errors")
	assert.Equal(t, len(usersList), 2, "should get 2 users")

	testEmails := []string{userA.Email, userB.Email}
	queriedEmails := make([]string, len(usersList))
	for i, u := range usersList {
		queriedEmails[i] = u.Email
	}

	for _, e := range testEmails {
		assert.Contains(t, queriedEmails, e, fmt.Sprintf("Email %s does not exists in queriedEmails %v", e, queriedEmails))
	}
}

func Test_CountSpecificUsers(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	userA := createUser(ctx, repo, "test1@test.com")
	_ = createUser(ctx, repo, "test2@test.com")

	count, err := repo.Count(ctx, fmt.Sprintf("email = \"%s\"", userA.Email))
	assert.NoError(t, err, "should count users without errors")
	assert.Equal(t, count, 1, "should count only 1 user")
}

func Test_CountAllUsers(t *testing.T) {
	repo := repositories.NewUserRepository()
	ctx := tests.TestContext()
	tests.ResetDatabase()

	_ = createUser(ctx, repo, "test1@test.com")
	_ = createUser(ctx, repo, "test2@test.com")

	count, err := repo.Count(ctx, "")
	assert.NoError(t, err, "should count users without errors")
	assert.Equal(t, count, 2, "should count 2 user")
}

func createUser(ctx context.Context, repo *repositories.UserRepository, email string) *models.User {
	h, _ := utils.CreatePasswordHash(email)
	user := models.User{
		Email:        email,
		PasswordHash: h,
	}
	_ = repo.Create(ctx, &user)

	return &user
}
