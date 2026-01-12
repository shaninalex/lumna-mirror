package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/config"
)

func NewController(conf *config.Config, router *gin.Engine) {
	router.POST("/auth/register", nil)
	router.POST("/auth/login", nil)
	router.POST("/auth/logout", nil)
	router.POST("/auth/refresh", nil)
	router.GET("/auth/oauth/github", nil)
	router.GET("/auth/oauth/github/callback", nil)
	router.GET("/auth/oauth/google", nil)
	router.GET("/auth/oauth/google/callback", nil)
}
