package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
	"gitlab.com/shaninalex/lumna/app/web/api/user"
)

func NewWebApplication(conf *config.Config) *gin.Engine {
	router := NewRouter()

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte(conf.SecretKey))))

	auth.NewController(conf, router)
	user.NewController(conf, router)

	return router
}
