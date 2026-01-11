package token

import (
	"net/http"
	"time"
)

const (
	// AccessTokenLifeTime is a 15 minutes duration
	AccessTokenLifeTime = 15 * time.Minute

	// RefreshTokenLifeTime is a 7 days duration
	RefreshTokenLifeTime = 7 * 24 * time.Hour

	// Issuer default issuer
	Issuer = "lumna"

	// NumericRefreshTokenLifeTime numeric refresh token lifetime ( 1 week )
	NumericRefreshTokenLifeTime = 7 * 24 * 60 * 60

	// NumericAccessTokenLifeTime numeric access token lifetime ( 15 min )
	NumericAccessTokenLifeTime = 15 * 60

	// Static auth cookie names

	// AccessTokenCookieName access token cookie name
	AccessTokenCookieName = "access_token"

	// RefreshTokenCookieName refresh token cookie name
	RefreshTokenCookieName = "refresh_token"
)

type AudToken string

var (
	AudTokenAPIUser   AudToken = "api"
	AudTokenAdminUser AudToken = "admin"
	AudTokenExternal  AudToken = "external"
)

// ClearAuthCookies clear all auth cookies from response to unauthenticate user
func ClearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // expire immediately
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // expire immediately
		SameSite: http.SameSiteStrictMode,
	})
}
