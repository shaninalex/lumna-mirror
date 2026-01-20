package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/projects"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
)

func NewApiRoutes(router *gin.RouterGroup) {
	auth.NewController(router)
	user.NewController(router)
	projects.NewController(router)
}
