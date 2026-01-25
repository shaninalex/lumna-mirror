package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.GET("/_health", utils.HealthRoute)
	router.Use(middlewares.CorsMiddleware())

	return router
}
