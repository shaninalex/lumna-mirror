package task

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (s *TaskController) Register(router *gin.RouterGroup) {
	router.POST("tasks", s.handleCreateTask)
	router.GET("tasks", s.handleListTask)
	router.GET("tasks/:taskId", s.handleGetTask)
	router.PATCH("tasks/:taskId", s.handlePatchTask)
}
