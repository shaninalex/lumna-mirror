package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key") // TODO: from config

func GenerateJWT(userID string, scopes []string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"scope": scopes,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
