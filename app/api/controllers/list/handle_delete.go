package list

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ListController) Delete(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.listService.Delete(c.Request.Context(), uint(listId)); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "List deleted")
}
