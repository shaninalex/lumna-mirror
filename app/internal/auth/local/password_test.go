package local_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
)

func Test_CreatePasswordHash(t *testing.T) {
	password := "mySecretPassword123!"

	hash, err := local.CreatePasswordHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash, "Hash should not be empty")
}

func Test_ValidatePasswordHash(t *testing.T) {
	password := "mySecretPassword123!"

	// Create hash
	hash, err := local.CreatePasswordHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Validate correct password
	err = local.ValidatePassword(hash, password)
	assert.NoError(t, err, "Password should validate successfully")

	// Validate wrong password
	err = local.ValidatePassword(hash, "wrongPassword")
	assert.Error(t, err, "Wrong password should return an error")
}

func Test_ShouldBeUnique(t *testing.T) {
	password := "mySecretPassword123!"
	hashA, _ := local.CreatePasswordHash(password)
	hashB, _ := local.CreatePasswordHash(password)

	assert.NotEqual(t, hashA, hashB, "Hashes should NOT be equal")
}
