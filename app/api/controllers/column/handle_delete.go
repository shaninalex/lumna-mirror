package column

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ColumnController) Delete(c *gin.Context) {
	columnId, err := strconv.Atoi(c.Param("columnId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.columnService.Delete(c.Request.Context(), uint(columnId)); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "Board deleted")
}
