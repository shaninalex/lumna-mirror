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

	var errs []error

	inv, _, err := s.invitationManager.Create(c.Request.Context(), data.Email, "member")
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		utils.Error(c, http.StatusBadRequest, errs[0])
		return
	}

	utils.Success(c, inv)
}
