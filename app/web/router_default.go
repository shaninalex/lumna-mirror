package web

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/web/api"
)

func NewDefaultRouter() *Router {
	r := NewRouter()

	r.GET("/_health", api.HealthRoute)

	return r
}

type RouterMiddleware interface {
	Wrap(next http.Handler) http.Handler
}

func (r *Router) ApplyMiddlewares(middlewares []RouterMiddleware) {
	for _, m := range middlewares {
		r.Use(m.Wrap)
	}
}
