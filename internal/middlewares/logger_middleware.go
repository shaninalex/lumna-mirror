// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/base"
)

type LoggerMiddleware struct {
	appName string
}

func NewLoggerMiddleware(name string) *LoggerMiddleware {
	return &LoggerMiddleware{appName: name}
}

func (m *LoggerMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(lrw, r.WithContext(context.WithValue(r.Context(), base.ContextAppName, m.appName)))

		duration := time.Since(start)
		log.Printf("[%s]: %s %s %d %v\n", m.appName, r.Method, r.URL.Path, lrw.status, duration)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}
