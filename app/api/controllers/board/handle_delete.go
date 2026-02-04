package board

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Delete(c *gin.Context) {
	boardId, err := uuid.Parse(c.Param("boardId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.boardService.BoardDelete(c.Request.Context(), boardId); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "Board deleted")
}

