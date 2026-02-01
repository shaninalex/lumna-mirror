package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/board"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/projects"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/internal/client"
)

func RegisterApiV1Routes(client *client.Client, baseRouter *gin.Engine) {
	// base API middlewares
	baseRouter.Use(middlewares.CORSMiddleware())

	router := baseRouter.Group("/api/v1/auth")
	auth.RegisterAuthController(router)

	privateRoutes := baseRouter.Group("/api/v1/")
	privateRoutes.Use(middlewares.AuthMiddleware)

	user.RegisterUserController(privateRoutes)
	projects.RegisterProjectsController(privateRoutes)
	column.RegisterColumnController(privateRoutes)
	board.RegisterBoardController(privateRoutes)
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

Or even:
Api.RegisterControllers([]controller{}, middlewares...)
for multiple controllers under single set of middlewares

*/
