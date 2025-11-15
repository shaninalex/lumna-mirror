package web

import (
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/pkg/base"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
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
