package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginClaims struct {
	jwt.RegisteredClaims
}

var loginJWTSecret = []byte("super-secret-key")

// GenerateLoginToken - generates token only for login. It's not access_token
func GenerateLoginToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"iss": "lumna-api",
		"aud": "lumna-web-client",
		"sub": userID,
		"typ": "login",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(loginJWTSecret))
}

// ValidateLoginToken - validates login_token
func ValidateLoginToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return loginJWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*LoginClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
