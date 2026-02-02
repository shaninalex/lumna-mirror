package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key") // TODO: from config

type AccessClaims struct {
	jwt.RegisteredClaims
	Scope string
}

// GenerateAccessJWTToken - generate jwt access token
func GenerateAccessJWTToken(userID string, scopes string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"iss":   "lumna-api",
		"aud":   "lumna-web-client",
		"sub":   userID,
		"scope": scopes,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseAccessJWTToken(tokenString string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to parse: %s", err)
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("access token expired")
	}

	return claims, nil
}
