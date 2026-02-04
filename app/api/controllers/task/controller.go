package task

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: services.NewTaskService(),
	}
}

func RegisterTaskController(router *gin.RouterGroup) {
	controller := NewTaskController()

	router.POST("task", controller.handleCreateTask)
}
