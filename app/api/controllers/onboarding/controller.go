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
	router.GET("state", s.workspaceExistsMiddleware(), s.handlerCheckState)
	router.POST("workspace", s.workspaceExistsMiddleware(), s.handlerWorkspace)
	router.POST("team", s.teammatesExistsMiddleware(), s.handlerTeam)
}
