package board

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

type BoardHandler struct {
	boardService *services.BoardService
	tasksService services.TaskManager
}

func NewBoardHandler() *BoardHandler {
	return &BoardHandler{
		boardService: services.NewBoardService(),
		tasksService: services.NewTaskManager(),
	}
}

func RegisterBoardController(router *web.Router) {
	h := NewBoardHandler()

	// Manage boards
	router.PATCH("/api/v1/board/{id}", h.Patch)
	router.DELETE("/api/v1/board/{id}", h.Delete)

	// Manage board lists
	router.GET("/api/v1/board/{id}/lists", h.ListsGet)
	router.POST("/api/v1/board/{id}/lists", h.ListsCreate)
	router.PATCH("/api/v1/lists/{listId}", h.ListsPatch)
	router.DELETE("/api/v1/lists/{listId}", h.ListsDelete)

	// Board tasks
	router.GET("/api/v1/board/{id}/tasks", h.TasksGet)
}
