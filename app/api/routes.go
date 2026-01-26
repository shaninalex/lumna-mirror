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
	// base API middlewares
	baseRouter.Use(middlewares.CORSMiddleware())

	// OAuth
	oauthRoutes := baseRouter.Group("/oauth")
	oauth.RegisterOAuthController(oauthRoutes)

	router := baseRouter.Group("/api/v1/auth")
	auth.RegisterAuthController(router)

	privateRoutes := baseRouter.Group("")

	// TODO: auth middleware

	user.RegisterUserController(privateRoutes)
	projects.RegisterProjectsController(privateRoutes)
}

/*

TODO: create single interface
Api.RegisterController(controller, middlewares...)

For example:
Api.RegisterController(
	NewProjectsController,
	AuthMiddleware(),
	PermissionsMiddleware(),
	TrackingMiddleware(),
)

*/
