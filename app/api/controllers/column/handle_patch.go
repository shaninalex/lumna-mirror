package column

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *ColumnController) Patch(c *gin.Context) {
	columnId, err := strconv.Atoi(c.Param("columnId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := services.ColumnUpdate{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column, err := s.columnService.Get(c.Request.Context(), uint(columnId))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column.Title = payload.Title

	if _, err := s.columnService.Update(c.Request.Context(), column); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, column)
}
