package onboarding

import (
	"github.com/gin-gonic/gin"
)

type OnboardingController struct {
}

func NewOnboardingController() *OnboardingController {
	s := &OnboardingController{}

	return s
}

func (s *OnboardingController) Register(router *gin.RouterGroup) {
	router.POST("workspace", s.handlerWorkspace)
	router.POST("team", s.handlerTeam)
}

func (s *OnboardingController) handlerWorkspace(c *gin.Context) {}
func (s *OnboardingController) handlerTeam(c *gin.Context)      {}
