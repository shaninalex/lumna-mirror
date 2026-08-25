package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TaskListQuery struct {
	BoardId uint `form:"board_id,omitempty"`
}

func (s *TaskController) handleListTask(c *gin.Context) {
	q := TaskListQuery{}
	if err := c.ShouldBindQuery(&q); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	tasks, boardTasks, err := s.taskService.List(c.Request.Context(), services.ServiceTaskListQuery{
		BoardId: q.BoardId,
	})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	result := adapters.ToTaskDtoList(tasks, boardTasks)

	utils.Success(c, result)
}
