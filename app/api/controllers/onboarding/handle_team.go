package onboarding

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

type TeamPage struct {
	Emails []string `json:"emails"`
}

func (s *OnboardingController) handlerTeam(c *gin.Context) {
	var data TeamPage
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	var errs []error

	wps, err := s.workspaceManager.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if len(wps) == 0 {
		utils.Error(c, http.StatusBadRequest, errors.New("no workspace created. Please repeat first step")) // Error codes instead
		return
	}

	var invitations []*models.Invitation
	for _, email := range data.Emails {
		inv, _, err := s.invitationManager.Create(c.Request.Context(), wps[0].ID, email, "member")
		if err != nil {
			errs = append(errs, err)
		}
		invitations = append(invitations, inv)
	}

	if len(errs) > 0 {
		utils.Error(c, http.StatusBadRequest, errs[0]) // TODO: handle this case
		return
	}

	utils.Success(c, invitations)
}
