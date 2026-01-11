package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
)

func NewController(conf *config.Config, router *gin.Engine) {
	router.GET("/user/me", nil)
}
