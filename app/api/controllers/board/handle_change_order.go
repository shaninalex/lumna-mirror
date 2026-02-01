package board

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) ChangeOrder(c *gin.Context) {

	// TODO: handle change
	// s.taskService.ReorderTask()
	// s.columnService.Reorder()

	utils.Success(c, nil)
}
