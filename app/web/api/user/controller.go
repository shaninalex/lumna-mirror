package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewController(router *gin.RouterGroup) {
	router.GET("/api/v1/user/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"email": "test@test.com",
		})
	})
}
