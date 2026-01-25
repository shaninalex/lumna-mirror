package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/projects"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
)

func RegisterApiV1Routes(conf *config.Config, baseRouter *gin.Engine) {
	router := baseRouter.Group("/api/v1")

	auth.NewController(router)

	privateRoutes := router.Group("")

	user.NewController(privateRoutes)
	projects.NewController(privateRoutes)
}
