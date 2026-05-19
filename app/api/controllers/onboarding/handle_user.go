package onboarding

import (
	"fmt"
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

	inv, token, err := s.invitationManager.Create(c.Request.Context(), data.Email, "member", map[string]any{
		"first_name": data.FirstName,
		"last_name":  data.LastName,
	})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	fmt.Println(inv)
	fmt.Println("Email link token: ", token)

	utils.Success(c, inv)
}
