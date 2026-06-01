package activity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/services/logger"
)

type ActivityController struct {
	logService logger.ActivityLogService
}

func NewActivityController(logService logger.ActivityLogService) *ActivityController {
	s := &ActivityController{
		logService: logService,
	}
	return s
}

func (s *ActivityController) Register(router *gin.RouterGroup) {
	router.GET("/activity/logs", s.handleLogsList)
}

func (s *ActivityController) handleLogsList(c *gin.Context) {
	entityType := c.Query("activity_type")
	entityIDStr := c.Query("activity_id")
	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	logs, err := s.logService.GetLog(c.Request.Context(), entityType, uint(entityID))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, logs)
}
