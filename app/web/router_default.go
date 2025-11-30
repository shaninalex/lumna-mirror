package web

import "gitlab.com/shaninalex/lumna/app/web/api"

func NewDefaultRouter() *Router {
	r := NewRouter()

	r.GET("/_health", api.HealthRoute)

	return r
}
