package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
)

type AuthController struct {
	localProvider *auth.EmailAuthProvider
}

func NewAuthContoller(authProvider *auth.EmailAuthProvider) *AuthController {
	s := &AuthController{
		localProvider: authProvider,
	}

	return s
}

func (s *AuthController) Register(router *gin.RouterGroup) {
	router.POST("/login", s.handleLogin)
	router.GET("/logout", s.handleLogout)

	// refresh require authenticated request to get user id
	router.GET("/refresh", s.handleRefresh)

	// router.GET("/api/v1/auth/oauth/github", nil)
	// router.GET("/api/v1/auth/oauth/github/callback", nil)

	// router.GET("/api/v1/auth/oauth/google", nil)
	// router.GET("/api/v1/auth/oauth/google/callback", nil)
}
