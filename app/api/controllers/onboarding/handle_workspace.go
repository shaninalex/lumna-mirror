package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

type WorkspacePage struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

func (s *OnboardingController) handlerWorkspace(c *gin.Context) {
	var data WorkspacePage
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	wp, err := s.workspaceManager.Create(c.Request.Context(), data.Email)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.workspaceManager.Update(c.Request.Context(), wp.ID, map[string]any{"owner_email": data.Email}); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	_, _, err = s.invitationManager.Create(c.Request.Context(), wp.ID, data.Email, "owner")
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, wp)
}
