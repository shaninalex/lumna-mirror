package column

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ColumnController struct {
	columnService *services.ColumnService
}

func NewColumnController(columnService *services.ColumnService) *ColumnController {
	s := &ColumnController{
		columnService: columnService,
	}

	return s
}

func (s *ColumnController) Register(router *gin.RouterGroup) {
	router.GET("columns", s.List)
	router.POST("columns", s.Create)
	router.DELETE("column/:columnId", s.Delete)
	router.PATCH("column/:columnId", s.Patch)
}
