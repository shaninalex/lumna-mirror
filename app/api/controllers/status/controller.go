package status

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type StatusController struct {
	statusService *services.StatusService
}

func NewStatusController(statusService *services.StatusService) *StatusController {
	s := &StatusController{
		statusService: statusService,
	}

	return s
}

func (s *StatusController) Register(router *gin.RouterGroup) {
	router.GET("statuses", s.List)
	router.POST("statuses", s.Create)
	router.DELETE("status/:statusId", s.Delete)
	router.PATCH("status/:statusId", s.Patch)
}
