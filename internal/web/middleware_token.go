// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/token"
)

type TokenMiddleware struct {
	accessTokenService token.AccessTokenService
}

func NewTokenMiddleware() *TokenMiddleware {
	return &TokenMiddleware{
		accessTokenService: token.NewDefaultAccessTokenService(),
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

		userID, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			//token.ClearAuthCookies(w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		// store claims in context
		ctx = context.WithValue(ctx, base.ContextUserID, uint(userID))

		// Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
