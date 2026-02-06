package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
)

func NewWebApplication() *gin.Engine {
	router := NewDefaultRouter()
	router.Use(middlewares.CORSMiddleware())
	api.RegisterApiV1Routes(router)

	return router
}
