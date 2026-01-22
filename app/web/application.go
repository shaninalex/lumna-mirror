package web

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
	"gitlab.com/shaninalex/lumna/app/web/pages"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func NewWebApplication(ctx context.Context) *gin.Engine {
	conf, ok := ctx.Value(internal.ContextConfig).(*config.Config)
	if !ok {
		panic("config not found in context")
	}

	router := NewRouter()

	// Middlewares
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ClientMiddleware(ctx))
	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))

	if conf.Html.Enabled {
		templRoutes := router.Group("/")
		{
			pages.NewPageRouter(templRoutes, conf)
		}
	}
	if conf.API.Enabled {
		apiRoutes := router.Group("/api/v1")
		{
			api.NewApiRoutes(apiRoutes)
		}
	}

	return router
}
