package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/oauth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/projects"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/internal/client"
)

func RegisterApiV1Routes(client *client.Client, baseRouter *gin.Engine) {
	oauthRoutes := baseRouter.Group("/oauth")
	oauth.RegisterOAuthController(oauthRoutes)

	router := baseRouter.Group("/api/v1")

	router.Use(middlewares.CORSMiddleware())

	auth.NewController(router)

	privateRoutes := router.Group("")

	user.NewController(privateRoutes)
	projects.NewController(privateRoutes)
}
