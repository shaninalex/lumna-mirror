// Copyright © 2025 Lumna. All rights reserved.

package web

import (
	"database/sql"

	"github.com/shaninalex/lumna/app/internal/base"
	"github.com/shaninalex/lumna/app/internal/db"
)

// DefaultRouter - default router.
func DefaultRouter(dbConnection *sql.DB) *Router {
	r := NewRouter()
	if base.IsDebug() {
		r.Use(NewRecoveryMiddleware().Wrap)
	}
	r.Use(db.NewMiddleware(dbConnection).Wrap)
	r.Use(NewLoggerMiddleware().Wrap)
	r.Use(NewCommonMiddleware().Wrap)
	r.Use(NewHeadersMiddleware().Wrap)
	r.GET("/_health", HandleHealth)

	return r
}
