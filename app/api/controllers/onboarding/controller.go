package onboarding

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type OnboardingController struct {
	workspaceManager  services.WorkspaceManager
	invitationManager services.InvitationManager
	user              *services.UserService
}

func NewOnboardingController(
	workspaceManager services.WorkspaceManager,
	invitationManager services.InvitationManager,
	user *services.UserService,
) *OnboardingController {
	s := &OnboardingController{
		workspaceManager:  workspaceManager,
		invitationManager: invitationManager,
		user:              user,
	}

	return s
}

func (s *OnboardingController) Register(router *gin.RouterGroup) {
	router.GET("state", s.handlerCheckState) // s.workspaceExistsMiddleware(), )
	router.POST("user", s.handlerUser)       // s.workspaceExistsMiddleware(), )
}
