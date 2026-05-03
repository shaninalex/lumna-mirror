package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argon2Time    = 1         // number of iterations
	argon2Memory  = 64 * 1024 // 64 MB
	argon2Threads = 4
	argon2KeyLen  = 32
	saltLen       = 16
)

// CreatePasswordHash generates a password hash
func CreatePasswordHash(rawPassword string) (string, error) {
	// Generate a random salt
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(rawPassword), salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argon2Memory, argon2Time, argon2Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encodedHash, nil
}

// ValidatePassword checks if the provided password matches the stored hash
func ValidatePassword(passwordHash, password string) error {
	parts := strings.Split(passwordHash, "$")
	if len(parts) != 6 {
		return errors.New("invalid hash format")
	}

	var memory, time, threads uint32
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, uint8(threads), uint32(len(expectedHash)))

	if subtleConstantTimeCompare(hash, expectedHash) {
		return nil
	}
	return errors.New("password does not match")
}

// subtleConstantTimeCompare compares two byte slices in constant time
func subtleConstantTimeCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
