package api

import (
	"net/http"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
)

// kratos hooks request headers:
// Authorization: <auth code>
// Content-Type: application/json
// Ory-Webhook-Request-Id: 5b1c5716-db2a-42b6-bc77-7b7619f73d86

// AuthHooksMiddleware - auth middleware.
type AuthHooksMiddleware struct {
	kratosService kratos.IKratos
	config        base.IConfig
}

// NewAuthHooksMiddleware - new auth hooks middleware.
func NewAuthHooksMiddleware(kratosService kratos.IKratos, config base.IConfig) *AuthHooksMiddleware {
	return &AuthHooksMiddleware{
		kratosService: kratosService,
		config:        config,
	}
}

// Wrap - wrapper, actual middleware
func (s *AuthHooksMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check Ory webhook request ID
		if r.Header.Get("Ory-Webhook-Request-Id") == "" {
			http.Error(w, "request id not set", http.StatusBadRequest)
			return
		}

		// Check Authorization
		expected := s.config.String("app.secret_key")
		if expected == "" || r.Header.Get("Authorization") != expected {
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		// additional validations here

		next.ServeHTTP(w, r)
	})
}
