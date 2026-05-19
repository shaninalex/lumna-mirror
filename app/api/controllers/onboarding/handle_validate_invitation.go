package onboarding

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *OnboardingController) ValidateInvitation(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.Error(c, http.StatusBadRequest, utils.NewApiError("invlid token - empty", "ERROR_INVITATION_INVALID_TOKEN", nil))
		return
	}

	if err := s.invitationManager.Validate(c.Request.Context(), token); err != nil {
		utils.Error(c, http.StatusBadRequest, utils.NewApiError(
			fmt.Sprintf("invlid token: %s", err.Error()),
			"ERROR_INVITATION_INVALID_TOKEN",
			nil,
		))
		return
	}

	invitation, err := s.invitationManager.Get(c.Request.Context(), token)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, utils.NewApiError(
			fmt.Sprintf("invlid token: %s", err.Error()),
			"ERROR_INVITATION_INVALID_TOKEN",
			nil,
		))
		return
	}

	utils.Success(c, invitation)
}
