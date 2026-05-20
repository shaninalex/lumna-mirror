package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateRefreshToken() (string, string, error) {
	b := make([]byte, 64)

	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}

	// to client
	plain := base64.RawURLEncoding.EncodeToString(b)

	// to store
	hash := ToHashToken(plain)

	return plain, hash, nil
}

func ToHashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
