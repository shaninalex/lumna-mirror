// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"net/http"

	"github.com/gorilla/csrf"
)

// CommonMiddleware - common middleware.
type CommonMiddleware struct {
}

// NewCommonMiddleware - new common middleware.
func NewCommonMiddleware() *CommonMiddleware {
	m := &CommonMiddleware{}
	return m
}

// Wrap - wrap.
func (m *CommonMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := readRequestIP(r)
		ctx := context.WithValue(r.Context(), "requestIP", ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func readRequestIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func CSRFMiddleware(secret []byte) Middleware {
	return func(next http.Handler) http.Handler {
		// double-submit cookie setup
		return csrf.Protect(
			secret,
			csrf.Secure(false), // true in prod (requires HTTPS)
			csrf.HttpOnly(true),
			csrf.Path("/"),
			csrf.CookieName("csrf_token"),
			csrf.FieldName("X-CSRF-Token"), // what we expect from frontend
		)(next)
	}
}
