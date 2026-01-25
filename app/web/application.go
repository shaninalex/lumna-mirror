package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/internal/client"
)

func NewWebApplication(client *client.Client) *gin.Engine {
	router := NewRouter()

	// Share application context to gin
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(client.Context())
		c.Next()
	})
	router.Use(middlewares.CORSMiddleware())

	if client.Config().API.Enabled {
		api.RegisterApiV1Routes(client, router)
	}

	return router
}
