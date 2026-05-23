package onboarding

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services"
)

type OnboardingController struct {
	workspaceManager  services.WorkspaceManager
	invitationManager services.InvitationManager
	userManager       services.UserManager
	config            *config.Config
}

func NewOnboardingController(
	workspaceManager services.WorkspaceManager,
	invitationManager services.InvitationManager,
	userManager services.UserManager,
	config *config.Config,
) *OnboardingController {
	s := &OnboardingController{
		workspaceManager:  workspaceManager,
		invitationManager: invitationManager,
		userManager:       userManager,
		config:            config,
	}

	return s
}

func (s *OnboardingController) Register(router *gin.RouterGroup) {
	router.POST("", s.handlerInit)
	router.GET("state", s.handlerCheckState)
	router.GET("invitation", s.ValidateInvitation)
}
