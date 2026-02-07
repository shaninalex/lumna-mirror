package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	s := &UserController{
		userService: userService,
	}

	return s
}

func (s *UserController) Register(router *gin.RouterGroup) {
	router.GET("/user/me", s.Me)
}
