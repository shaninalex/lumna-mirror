package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func generateSecureCode() (string, error) {
	const size = 32 // 32 bytes = 256 bits
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func TestGenerateSecureCode(t *testing.T) {
	code1, _ := generateSecureCode()
	code2, _ := generateSecureCode()

	if code1 == code2 {
		t.Fatal("codes should be unique")
	}
}
