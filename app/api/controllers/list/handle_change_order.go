package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *ListController) ChangeOrder(c *gin.Context) {
	var payload services.KanbanListChangeOrderPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.listService.Reorder(c.Request.Context(), &payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if payload.Activity != nil {
		if id, err := utils.GetUserID(c); err == nil {
			s.activityLogger.Log(id, payload.Activity)
		}
	}

	utils.Success(c, nil)
}
