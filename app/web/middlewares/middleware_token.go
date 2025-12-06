package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/shaninalex/lumna/app/global"
	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TokenMiddleware struct {
	accessTokenService services.AccessTokenService
}

func NewTokenMiddleware() *TokenMiddleware {
	return &TokenMiddleware{
		accessTokenService: services.NewDefaultAccessTokenService(),
	}
}

func (s *TokenMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to read token from Authorization header first
		tokenString := ""
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
				tokenString = parts[1]
			}
		}

		// If not in header, try HTTP-only cookie
		if tokenString == "" {
			cookie, err := r.Cookie(token.AccessTokenCookieName)
			if err != nil {
				//token.ClearAuthCookies(w)
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}
			tokenString = cookie.Value
		}

		// validate
		claims, err := s.accessTokenService.Validate(tokenString, token.AudTokenAPIUser)
		if err != nil {
			//token.ClearAuthCookies(w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			//token.ClearAuthCookies(w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		// store claims in context
		ctx = context.WithValue(ctx, global.ContextUserID, userID)

		// Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
