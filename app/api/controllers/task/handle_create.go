package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) handleCreateTask(c *gin.Context) {
	payload := services.TaskPayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	task, err := s.taskService.CreateTask(c.Request.Context(), &payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	utils.Success(c, task)
}
