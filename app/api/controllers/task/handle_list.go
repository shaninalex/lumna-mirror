package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) handleListTask(c *gin.Context) {
	listQuery := services.ServiceTaskListQuery{}
	if err := c.ShouldBindQuery(&listQuery); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	tasks, err := s.taskService.List(c.Request.Context(), listQuery)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, tasks)
}
