// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"database/sql"
	"net/http"
	"path"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web/middlewares"
	"gorm.io/gorm"
)

// DefaultRouter - default router.
func DefaultRouter(db *gorm.DB, name string) *fiber.App {
	router := fiber.New(fiber.Config{
		AppName: name,
	})
	router.Use(recover.New())
	router.Use(logger.New())
	router.Use(database.NewDbMiddleware(db))
	router.Use(NewCommonMiddleware())
	router.Get("/_health", HandleHealth)
	return router
}

// AuthRouter - auth router.
func AuthRouter(cnf base.IConfig, db *gorm.DB, name string) *fiber.App {
	router := DefaultRouter(db, name)
	kratosClient := kratos.NewKratosService(cnf.String("kratos.url_browser"))
	router.Use(NewAuthMiddleware(kratosClient))
	return router
}

type Middleware func(http.Handler) http.Handler

type Router struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// NewTestRouter creates a new Router for testing purposes without metrics.
func NewTestRouter(*testing.T) *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.HandlerFunc("GET", path, handler)
}

func (r *Router) HEAD(path string, handler http.HandlerFunc) {
	r.HandlerFunc("HEAD", path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.HandlerFunc("POST", path, handler)
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.HandlerFunc("PUT", path, handler)
}

func (r *Router) PATCH(path string, handler http.HandlerFunc) {
	r.HandlerFunc("PATCH", path, handler)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.HandlerFunc("DELETE", path, handler)
}

func (r *Router) HandlerFunc(method, route string, handler http.HandlerFunc) {
	for _, pattern := range []string{
		method + " " + path.Join(route),
		method + " " + path.Join(route, "{$}"),
	} {
		r.handleWithAllMiddlewares(r.mux, pattern, handler)
	}
}

func (r *Router) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) handleWithAllMiddlewares(mux *http.ServeMux, pattern string, handler http.Handler) {
	// Apply global middlewares
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		//NoCache(w)
		handler.ServeHTTP(w, req)
	})
}

//func NoCache(w http.ResponseWriter) {
//	w.Header().Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
//}

func NewAppRouter(db *sql.DB, appName string) *Router {
	r := NewRouter()
	r.Use(middlewares.NewRecoveryMiddleware().Wrap)
	r.Use(database.NewMiddleware(db).Wrap)
	r.Use(middlewares.NewLoggerMiddleware(appName).Wrap)
	return r
}
