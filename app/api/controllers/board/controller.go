package board

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/logger"
	"gitlab.com/shaninalex/lumna/app/services"
)

type BoardController struct {
	boardService    *services.BoardService
	columnService   *services.ColumnService
	taskService     *services.TaskService
	activityService *logger.ActivityService
}

func NewBoardController(
	boardService *services.BoardService,
	columnService *services.ColumnService,
	taskService *services.TaskService,
	activityService *logger.ActivityService,
) *BoardController {
	s := &BoardController{
		boardService:    boardService,
		columnService:   columnService,
		taskService:     taskService,
		activityService: activityService,
	}
	return s
}

func (s *BoardController) Register(router *gin.RouterGroup) {
	router.POST("boards", s.Create)
	router.GET("board/:boardId", s.Get)
	router.PATCH("board/:boardId", s.Patch)
	router.DELETE("board/:boardId", s.Delete)
	router.PATCH("board/:boardId/order", s.ChangeOrder)
}
