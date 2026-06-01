package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/services/user"
)

type UserController struct {
	userService       user.UserManager
	invitationManager services.InvitationManager
}

func NewUserController(
	userService user.UserManager,
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
