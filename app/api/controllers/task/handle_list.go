package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *TaskController) handleListTask(c *gin.Context) {
	boardId, err := strconv.Atoi(c.Query("board_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	tasks, err := s.taskService.List(c.Request.Context(), map[string]any{"board_id": boardId})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, tasks)
}
