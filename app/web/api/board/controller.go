package board

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

type BoardHandler struct {
	boardService *services.BoardService
}

func NewProjectHandler() *BoardHandler {
	return &BoardHandler{
		boardService: services.NewBoardService(),
	}
}

func RegisterProjectController(router *web.Router) {
	h := NewProjectHandler()

	// Manage boards
	router.PATCH("/api/v1/board/{id}/lists", h.Patch)
	router.DELETE("/api/v1/board/{id}/lists", h.Delete)

	// Manage board lists
	router.GET("/api/v1/board/{id}/lists", h.ListsGet)
	router.POST("/api/v1/board/{id}/lists", h.ListsCreate)
	router.PATCH("/api/v1/lists/{listId}", h.ListsPatch)
	router.DELETE("/api/v1/lists/{listId}", h.ListsDelete)
}
