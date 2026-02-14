package column

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

var (
	ErrorNoBoardProvided = errors.New("no board_id provided or it's invalid")
)

func (s *ColumnController) List(c *gin.Context) {
	boardId, err := strconv.Atoi(c.Query("board_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, ErrorNoBoardProvided)
		return
	}

	columns := s.columnService.Filter(c.Request.Context(), uint(boardId))
	utils.Success(c, columns)
}
