package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
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

	tasks, err := s.taskService.List(c.Request.Context(), services.ServiceTaskListQuery{
		BoardId: q.BoardId,
	})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	ids := make([]int64, len(tasks))
	for i, t := range tasks {
		ids[i] = t.ID
	}

	events := []models.EntityEvent{}
	events, _ = s.entityEventService.ListByEntityIds(c.Request.Context(), ids, "task")
	result := adapters.ToTaskDtoList(tasks, events)

	utils.Success(c, result)
}
