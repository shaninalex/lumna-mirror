package board

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Patch(c *gin.Context) {
	boardId, err := uuid.Parse(c.Param("boardId"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := struct {
		ProjectId uuid.UUID `json:"project_id"`
		Title     string    `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	board, err := s.boardService.Get(c.Request.Context(), boardId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	board.Title = payload.Title

	if err := s.boardService.Update(c.Request.Context(), board); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, board)
}
