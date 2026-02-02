package column

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ColumnController struct {
	columnService *services.ColumnService
}

func NewColumnController() *ColumnController {
	s := &ColumnController{
		columnService: services.NewColumnService(),
	}

	return s
}

func RegisterColumnController(router *gin.RouterGroup) {
	controller := NewColumnController()

	router.GET("columns", controller.List)
	router.POST("columns", controller.Create)
	router.DELETE("column/:columnId", controller.Delete)
	router.PATCH("column/:columnId", controller.Patch)
}
