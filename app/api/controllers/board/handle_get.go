package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Get(c *gin.Context) {
	boardId, err := strconv.Atoi(c.Param("boardId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	board, err := s.boardService.Get(c.Request.Context(), uint(boardId))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, board)
}
