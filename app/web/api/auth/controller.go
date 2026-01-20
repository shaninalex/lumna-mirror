package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
)

type AuthController struct {
	localProvider *local.LocalAuthProvider
}

func NewAuthContoller() *AuthController {
	s := &AuthController{
		localProvider: local.NewLocalAuthProvider(),
	}

	return s
}

func NewController(router *gin.RouterGroup) {

	controller := &AuthController{}

	// router.POST("/auth/register", nil)
	router.POST("/auth/login", controller.HandleAuthLogin)
	// router.POST("/api/v1/auth/logout", nil)
	// router.POST("/api/v1/auth/refresh", nil)

	// router.GET("/api/v1/auth/oauth/github", nil)
	// router.GET("/api/v1/auth/oauth/github/callback", nil)

	// router.GET("/api/v1/auth/oauth/google", nil)
	// router.GET("/api/v1/auth/oauth/google/callback", nil)
}
