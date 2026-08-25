package task

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *TaskController) handleCreateTask(c *gin.Context) {
	payload := models.TaskCreateOnBoard{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	b, _ := json.Marshal(payload)
	fmt.Println(string(b))

	task, err := s.taskService.CreateTask(c.Request.Context(), &payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	taskDto := adapters.ToTaskDto(*task, models.BoardTask{
		BoardId:  *payload.BoardId,
		TaskId:   task.ID,
		ColumnId: *payload.ColumnId,
		Position: *payload.Position,
	})

	utils.Success(c, taskDto)
}
