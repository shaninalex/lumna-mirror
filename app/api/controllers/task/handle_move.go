package task

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) handleMoveTask(c *gin.Context) {
	var data dto.KanbanMoveTaskDto
	if err := c.BindJSON(&data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	err := s.taskService.Move(c.Request.Context(), int64(data.BoardId), models.RearangeTask{
		ColumnId: data.ColumnId,
		Tasks:    data.Tasks,
	})
	if err != nil {
		utils.Error(c, 400, err)
		return
	}

	tasks, err := s.taskService.List(c.Request.Context(), services.ServiceTaskListQuery{
		BoardId: uint(data.BoardId),
	})

	if err != nil {
		utils.Error(c, 400, err)
		return
	}

	utils.Success(c, tasks)
}
