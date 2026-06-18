package sprint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *SprintController) handleGetSprint(c *gin.Context) {
	sprintId, err := strconv.Atoi(c.Param("sprintId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	sprint, err := s.sprintService.GetByID(c.Request.Context(), uint(sprintId))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, sprint)
}
