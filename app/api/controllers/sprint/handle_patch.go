package sprint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *SprintController) handlePatchSprint(c *gin.Context) {
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

	payload := models.Sprint{}
	if err = c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	sprint.Title = payload.Title
	sprint.Description = payload.Description
	sprint.Done = payload.Done
	sprint.ProjectID = payload.ProjectID
	sprint.StartedAt = payload.StartedAt
	sprint.FinishedAt = payload.FinishedAt

	if err = s.sprintService.Update(c.Request.Context(), sprint); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, sprint)
}
