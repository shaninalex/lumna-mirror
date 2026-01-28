package oauth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func generateSecureCode() (string, error) {
	const size = 32 // 32 bytes = 256 bits
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func generateRefreshToken() (string, string, error) {
	b := make([]byte, 64)

	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}

	// to client
	plain := base64.RawURLEncoding.EncodeToString(b)

	// to store
	hash := hashToken(plain)

	return plain, hash, nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
