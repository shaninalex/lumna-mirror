package web

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/static"
	"gitlab.com/shaninalex/lumna/app/web/api"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func NewWebApplication(ctx context.Context) *gin.Engine {
	conf, ok := ctx.Value(internal.ContextConfig).(*config.Config)
	if !ok {
		panic("config not found in context")
	}

	router := NewRouter()

	// Middlewares
	// router.Use(middlewares.Logger())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ClientMiddleware(ctx))
	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))

	// Controllers
	if staticFS := static.GetStaticFS(); staticFS != nil {
		router.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			static.SPAHandler(staticFS)(c)
		})
	}

	apiRoutes := router.Group("/api")
	{
		auth.NewController(apiRoutes)
		apiRoutes.GET("/_health", api.HealthRoute)
	}

	privateApiRoutes := router.Group("/api")
	privateApiRoutes.Use(middlewares.AuthRequired())
	{
		user.NewController(privateApiRoutes)
	}

	return router
}
