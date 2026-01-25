package oauth

import (
	"crypto/sha256"
	"encoding/base64"
	"slices"
)

func isRedirectAllowed(allowed []string, uri string) bool {
	return slices.Contains(allowed, uri)
}

func validatePKCE(challenge, method string) bool {
	if challenge == "" {
		return false
	}
	if method == "" {
		method = "S256"
	}
	return method == "S256"
}

func verifyPKCE(verifier, expectedChallenge string) bool {
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])
	return challenge == expectedChallenge
}
