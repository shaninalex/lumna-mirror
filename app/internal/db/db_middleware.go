package db

import (
	"context"
	"database/sql"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/internal/base"
)

type Middleware struct {
	db *sql.DB
}

func NewMiddleware(db *sql.DB) *Middleware {
	return &Middleware{db: db}
}

func (s *Middleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), base.ContextDB, s.db))
		next.ServeHTTP(w, r)
	})
}
