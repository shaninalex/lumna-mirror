package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func Test_CreatePasswordHash(t *testing.T) {
	password := "mySecretPassword123!"

	hash, err := utils.CreatePasswordHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash, "Hash should not be empty")
}

func Test_ValidatePasswordHash(t *testing.T) {
	password := "mySecretPassword123!"

	// Create hash
	hash, err := utils.CreatePasswordHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Validate correct password
	err = utils.ValidatePassword(hash, password)
	assert.NoError(t, err, "Password should validate successfully")

	// Validate wrong password
	err = utils.ValidatePassword(hash, "wrongPassword")
	assert.Error(t, err, "Wrong password should return an error")
}

func Test_ShouldBeUnique(t *testing.T) {
	password := "mySecretPassword123!"
	hashA, _ := utils.CreatePasswordHash(password)
	hashB, _ := utils.CreatePasswordHash(password)

	assert.NotEqual(t, hashA, hashB, "Hashes should NOT be equal")
}
