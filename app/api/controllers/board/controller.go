package board

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type BoardController struct {
	boardService  *services.BoardService
	columnService *services.ColumnService
	taskService   *services.TaskService
}

func NewBoardController() *BoardController {
	s := &BoardController{
		boardService:  services.NewBoardService(),
		columnService: services.NewColumnService(),
		taskService:   services.NewTaskService(),
	}
	return s
}

func RegisterBoardController(router *gin.RouterGroup) {
	controller := NewBoardController()

	router.POST("boards", controller.Create)
	router.PATCH("board/:boardId", controller.Patch)
	router.DELETE("board/:boardId", controller.Delete)
	router.PATCH("board/:boardId/order", controller.ChangeOrder)

}
