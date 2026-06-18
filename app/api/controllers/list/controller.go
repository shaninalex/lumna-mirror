package list

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/services/logger"
)

type ListController struct {
	listService    services.ListService
	columnService  services.StatusService
	taskService    services.TaskService
	activityLogger logger.ActivityLogger
}

func NewListController(
	listService services.ListService,
	columnService services.StatusService,
	taskService services.TaskService,
	activityLogger logger.ActivityLogger,
) *ListController {
	s := &ListController{
		listService:    listService,
		columnService:  columnService,
		taskService:    taskService,
		activityLogger: activityLogger,
	}
	return s
}

func (s *ListController) Register(router *gin.RouterGroup) {
	router.POST("lists", s.Create)
	router.GET("lists", s.List)
	router.GET("lists/:listId", s.Get)
	router.PATCH("lists/:listId", s.Patch)
	router.DELETE("lists/:listId", s.Delete)
	router.PATCH("lists/:listId/order", s.ChangeOrder)
}
