package list

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ListController struct {
	listService   *services.ListService
	columnService  *services.StatusService
	taskService    *services.TaskService
	activityLogger *logger.ActivityLogger
}

func NewListController(
	listService *services.ListService,
	columnService *services.StatusService,
	taskService *services.TaskService,
	activityLogger *logger.ActivityLogger,
) *ListController {
	s := &ListController{
		listService:   listService,
		columnService:  columnService,
		taskService:    taskService,
		activityLogger: activityLogger,
	}
	return s
}

func (s *ListController) Register(router *gin.RouterGroup) {
	router.POST("lists", s.Create)
	router.GET("lists", s.List)
	router.GET("list/:listId", s.Get)
	router.PATCH("list/:listId", s.Patch)
	router.DELETE("list/:listId", s.Delete)
	router.PATCH("list/:listId/order", s.ChangeOrder)
}
