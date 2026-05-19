package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

type OnboardingState string

var (
	OnboardingStateDisabled  OnboardingState = "DISABLED"
	OnboardingStateAllowed   OnboardingState = "ALLOWED"
	OnboardingStateCompleted OnboardingState = "COMPLETED"
)

func (s *OnboardingController) handlerCheckState(c *gin.Context) {
	if s.config.Bool("onboarding.allowed") {
		utils.Success(c, map[string]any{
			"state": OnboardingStateDisabled,
		})
	}

	users, err := s.userManager.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}
	if len(users) == 0 {
		// No users exists in application, onboarding allowed.
		utils.Success(c, map[string]any{
			"state": OnboardingStateAllowed,
		})
		return
	}

	utils.Success(c, map[string]any{
		"state": OnboardingStateCompleted,
	})
}
