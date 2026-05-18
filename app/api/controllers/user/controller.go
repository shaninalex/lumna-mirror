package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type UserController struct {
	userService       services.UserManager
	invitationManager services.InvitationManager
}

func NewUserController(
	userService services.UserManager,
	invitationManager services.InvitationManager,
) *UserController {
	s := &UserController{
		userService:       userService,
		invitationManager: invitationManager,
	}

	return s
}

func (s *UserController) Register(router *gin.RouterGroup) {
	router.GET("/user/me", s.Me)
}
