package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
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

	utils.Success(c, data)
}
