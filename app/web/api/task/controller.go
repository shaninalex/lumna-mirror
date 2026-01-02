package task

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

type TaskHandler struct {
	tasksService services.TaskManager
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{
		tasksService: services.NewTaskManager(),
	}
}

func RegisterTaskController(router *web.Router) {
	h := NewTaskHandler()

	// Manage tasks
	router.GET("/api/v1/task/{id}", h.Get)
	router.PATCH("/api/v1/task/{id}", h.Patch)
	router.DELETE("/api/v1/task/{id}", h.Delete)
}
