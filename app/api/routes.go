package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/board"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/project"
	"gitlab.com/shaninalex/lumna/app/api/controllers/task"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
)

func RegisterApiV1Routes(baseRouter *gin.Engine) {
	// base API middlewares
	baseRouter.Use(middlewares.CORSMiddleware())

	router := baseRouter.Group("/api/v1/auth")
	auth.RegisterAuthController(router)

	privateRoutes := baseRouter.Group("/api/v1/")
	privateRoutes.Use(middlewares.AuthMiddleware)

	user.RegisterUserController(privateRoutes)
	project.RegisterProjectController(privateRoutes)
	column.RegisterColumnController(privateRoutes)
	board.RegisterBoardController(privateRoutes)
	task.RegisterTaskController(privateRoutes)
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
