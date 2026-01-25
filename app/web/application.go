package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/web/api"
)

func NewWebApplication(client *client.Client) *gin.Engine {
	conf := client.Config()
	router := NewRouter()

	// Share application context to gin
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(client.Context())
		c.Next()
	})

	if conf.API.Enabled {
		api.RegisterApiV1Routes(conf, router)
	}

	return router
}
