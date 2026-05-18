package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

type OnboardingUserPage struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *OnboardingController) handlerUser(c *gin.Context) {
	var data OnboardingUserPage
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	inv, _, err := s.invitationManager.Create(c.Request.Context(), data.Email, "member", nil)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, inv)
}
