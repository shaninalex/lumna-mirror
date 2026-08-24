package column

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *ColumnController) Patch(c *gin.Context) {
	statusId, err := strconv.Atoi(c.Param("status_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := services.BoardUpdate{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	status, err := s.columnService.Get(c.Request.Context(), uint(statusId))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	status.Title = payload.Title

	if _, err := s.columnService.Update(c.Request.Context(), status); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, status)
}
