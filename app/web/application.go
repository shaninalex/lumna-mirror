package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
)

func NewWebApplication(conf *config.Config) *gin.Engine {
	router := NewRouter()

	auth.NewController(conf, router)
	user.NewController(conf, router)

	return router
}
