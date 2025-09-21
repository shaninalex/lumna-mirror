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
		tokenString := r.Header.Get("Authorization")
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenService := new(TokenService)
		claims, err := tokenService.ValidateToken(r.Context(), parts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "jti", claims["jti"].(string))
		ctx = context.WithValue(ctx, base.ContextUserID, claims["sub"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
