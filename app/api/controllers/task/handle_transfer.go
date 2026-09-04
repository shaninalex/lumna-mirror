package task

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) handletransferTask(c *gin.Context) {
	var data models.TransferTaskBetweenColumns
	if err := c.BindJSON(&data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	if err := s.taskService.Transfer(c.Request.Context(), data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	// NOTE: ideally would be great just return all the tasks by id. But it does not metter for now
	tasks, err := s.taskService.List(c.Request.Context(), services.ServiceTaskListQuery{
		BoardId: uint(data.BoardId),
	})
	if err != nil {
		utils.Error(c, 400, err)
		return
	}

	utils.Success(c, tasks)
}
