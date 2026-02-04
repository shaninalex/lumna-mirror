package column

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ColumnController) Delete(c *gin.Context) {
	columnId, err := uuid.Parse(c.Param("columnId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.columnService.Delete(c.Request.Context(), columnId); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "Board deleted")
}
