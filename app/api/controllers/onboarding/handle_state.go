package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *OnboardingController) handlerCheckState(c *gin.Context) {
	wp, err := s.workspaceManager.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}
	if len(wp) == 0 {
		utils.Success(c, map[string]any{
			"state": "WORKSPACES",
		})
		return
	}

	users, err := s.userManager.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}
	if len(users) == 0 {
		// That mean that workspace was created but owner do not accept invite and do not create a user
		utils.Success(c, map[string]any{
			"state": "COMPLETED",
		})
		return
	}

	utils.Success(c, map[string]any{
		"state": "COMPLETED",
	})
}
