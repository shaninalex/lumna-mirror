// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"net/http"
	"strings"

	"gitlab.com/shaninalex/flowreon/internal/base"
)

func TokenMiddleware(next http.Handler) http.Handler {
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
			cookie, err := r.Cookie("access_token")
			if err != nil {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}
			tokenString = cookie.Value
		}

		// validate
		tokenService := new(TokenService)
		claims, err := tokenService.ValidateToken(r.Context(), tokenString)
		if err != nil {
			ClearAccessTokenCookie(w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// store claims in context
		ctx := context.WithValue(r.Context(), "jti", uint(claims["jti"].(float64)))
		ctx = context.WithValue(ctx, base.ContextUserID, uint(claims["sub"].(float64)))

		// Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
