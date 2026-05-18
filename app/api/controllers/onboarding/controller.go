package onboarding

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type OnboardingController struct {
	workspaceManager  services.WorkspaceManager
	invitationManager services.InvitationManager
	userManager       services.UserManager
}

func NewOnboardingController(
	workspaceManager services.WorkspaceManager,
	invitationManager services.InvitationManager,
	userManager services.UserManager,
) *OnboardingController {
	s := &OnboardingController{
		workspaceManager:  workspaceManager,
		invitationManager: invitationManager,
		userManager:       userManager,
	}

	return s
}

func (s *OnboardingController) Register(router *gin.RouterGroup) {
	router.GET("state", s.handlerCheckState)
	router.POST("user", s.handlerUser)
	router.GET("invitation", s.ValidateInvitation)
}
