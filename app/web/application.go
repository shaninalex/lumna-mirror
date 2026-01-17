package web

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
)

func NewWebApplication(ctx context.Context) *gin.Engine {
	conf, ok := ctx.Value(internal.ContextConfig).(*config.Config)
	if !ok {
		panic("config not found in context")
	}

	router := NewRouter()

	// Middlewares
	// router.Use(middlewares.Logger())
	router.Use(cors.Default())
	router.Use(middlewares.ClientMiddleware(ctx))
	router.Use(sessions.Sessions("session", NewCookieStore(conf)))

	// Controllers
	public := router.Group("/")
	auth.NewController(public)

	private := router.Group("/")
	private.Use(middlewares.AuthRequired())
	user.NewController(private)

	return router
}
