package sprint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *SprintController) handleCreateSprint(c *gin.Context) {
	sprint := models.Sprint{}

	if err := c.ShouldBindJSON(&sprint); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	err := s.sprintService.Create(c.Request.Context(), &sprint)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, sprint)
}
