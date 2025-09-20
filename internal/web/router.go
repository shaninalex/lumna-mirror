// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"database/sql"
	"net/http"
	"path"

	"github.com/gorilla/csrf"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/database"
)

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

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	handler.ServeHTTP(w, req)
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
		r.mux.HandleFunc(pattern, handler)
	}
}

func (r *Router) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) handleWithAllMiddlewares(mux *http.ServeMux, pattern string, handler http.Handler) {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		//NoCache(w)
		handler.ServeHTTP(w, req)
	})
}

func (r *Router) Run() error {
	csrfMiddleware := csrf.Protect(
		[]byte("32-byte-long-auth-key-123456789012"),
		csrf.Secure(false), // requires HTTPS. True in production
		csrf.Path("/"),
		csrf.FieldName("csrf_token"),
		csrf.CookieName("csrf_token"),
		csrf.TrustedOrigins([]string{"localhost:8000", "127.0.0.1:8000"}),
	)
	return http.ListenAndServe(":8000", csrfMiddleware(r))
}

// DefaultRouter - default router.
func DefaultRouter(db *sql.DB, name string) *Router {
	r := NewRouter()
	r.Use(NewRecoveryMiddleware().Wrap)
	r.Use(database.NewMiddleware(db).Wrap)
	r.Use(NewLoggerMiddleware(name).Wrap)
	r.Use(NewCommonMiddleware().Wrap)
	r.GET("/_health", HandleHealth)
	return r
}

// AuthRouter - auth router.
func AuthRouter(cnf base.IConfig, db *sql.DB, name string) *Router {
	router := DefaultRouter(db, name)
	//kratosClient := kratos.NewKratosService(cnf.String("kratos.url_browser"))
	//router.Use(NewAuthMiddleware(kratosClient).Wrap)
	return router
}
