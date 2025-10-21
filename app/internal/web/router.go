// Copyright © 2025 Lumna. All rights reserved.

package web

import (
	"database/sql"
	"log"
	"net/http"
	"path"

	"github.com/shaninalex/lumna/app/internal/base"
	"github.com/shaninalex/lumna/app/internal/db"
)

type Middleware func(http.Handler) http.Handler
type route struct {
	method string
	path   string
}

type Router struct {
	mux         *http.ServeMux
	middlewares []Middleware
	routes      []route
}

func NewRouter() *Router {
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

func (r *Router) HandlerFunc(method, url string, handler http.HandlerFunc) {
	r.routes = append(r.routes, route{method: method, path: url})

	for _, pattern := range []string{
		method + " " + path.Join(url),
		method + " " + path.Join(url, "{$}"),
	} {
		r.handleWithAllMiddlewares(r.mux, pattern, handler)
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
	//if base.IsDebug() {
	//	r.printRoutes()
	//}
	log.Println("server started... on port :8000")
	return http.ListenAndServe(":8000", corsMiddleware(r))
}

//func (r *Router) printRoutes() {
//	for _, rt := range r.routes {
//		log.Printf("%s %s\n", rt.method, rt.path)
//	}
//}

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
