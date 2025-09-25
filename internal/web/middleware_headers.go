// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// HeadersMiddleware - headers middleware.
type HeadersMiddleware struct {
}

// NewHeadersMiddleware - new headers middleware.
func NewHeadersMiddleware() *HeadersMiddleware {
	m := &HeadersMiddleware{}
	return m
}

// Wrap - wrap.
func (m *HeadersMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		elapsed := time.Since(start).Milliseconds()
		w.Header().Set("X-Response-Time", strconv.FormatInt(elapsed, 10)+"ms")
		w.Header().Set("X-API-Version", "1")
		w.Header().Set("X-Request-ID", uuid.NewString())
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	})
}
