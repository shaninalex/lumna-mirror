package web

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/projects"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
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
	// router.Use(middlewares.Logger())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ClientMiddleware(ctx))
	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))

	// API Controllers
	apiRoutes := router.Group("/api")
	{
		auth.NewController(apiRoutes)
		apiRoutes.GET("/_health", api.HealthRoute)
	}

	privateApiRoutes := router.Group("/api/v1")
	privateApiRoutes.Use(middlewares.AuthRequired())
	{
		user.NewController(privateApiRoutes)
		projects.NewController(privateApiRoutes)
	}

	// Conditionally apply embedded web client build
	// if staticFS := static.GetStaticFS(); staticFS != nil {
	// 	log.Println("[Lumna] embedd static files")
	// 	spaHandler := static.SPAHandler(staticFS)
	// 	router.NoRoute(func(c *gin.Context) {
	// 		if strings.HasPrefix(c.Request.URL.Path, "/api") {
	// 			utils.Error(c, http.StatusNotFound, fmt.Errorf("not found"))
	// 			return
	// 		}
	// 		spaHandler(c)
	// 	})
	// } else {
	templRoutes := router.Group("/")
	{
		pages.NewPageRouter(templRoutes, conf)
	}
	//
	// }

	return router
}
