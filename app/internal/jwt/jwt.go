package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key") // TODO: from config

// GenerateAccessJWTToken - generate jwt access token
func GenerateAccessJWTToken(userID string, scopes []string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"scope": scopes,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// TODO: use separate jwtSecret for login_token

var loginJWTSecret = []byte("super-secret-key")

// GenerateLoginToken - generates token only for login. It's not access_token
func GenerateLoginToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"iss": "https://auth.lumna.local",
		"sub": userID,
		"typ": "login",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(loginJWTSecret))
}

// ValidateLoginToken - validates login_token
func ValidateLoginToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return loginJWTSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	return claims.Subject, nil
}
