package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *TaskController) handleListTask(c *gin.Context) {
	q, err := adapters.ParseQueryParams(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if q == nil {
		utils.Error(c, http.StatusBadRequest, adapters.TaskListParseParamsError)
		return
	}

	tasks, err := s.taskService.List(c.Request.Context(), *q)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	ids := make([]int64, len(tasks))
	for i, t := range tasks {
		ids[i] = t.ID
	}

	events := []models.EntityEvent{}
	events, err = s.entityEventService.ListByEntityIds(c.Request.Context(), ids, "task")
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	result := adapters.ToTaskDtoList(tasks, events)
	utils.Success(c, result)
}
