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

	utils.Success(c, data)
}
