package sprint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

func (s *SprintController) handleListSprint(c *gin.Context) {
	listQuery := repositories.SprintListQuery{}
	if err := c.ShouldBindQuery(&listQuery); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	sprints, err := s.sprintService.List(c.Request.Context(), listQuery)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, sprints)
}
