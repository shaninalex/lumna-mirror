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

func RegisterAuthController(router *gin.RouterGroup) {
	controller := &AuthController{}

	router.POST("/login", controller.handleLogin)
	router.GET("/logout", controller.handleLogout)
	router.GET("/refresh", controller.handleRefresh)

	// router.GET("/api/v1/auth/oauth/github", nil)
	// router.GET("/api/v1/auth/oauth/github/callback", nil)

	// router.GET("/api/v1/auth/oauth/google", nil)
	// router.GET("/api/v1/auth/oauth/google/callback", nil)
}
