package sprint

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services/sprint"
)

type SprintController struct {
	sprintService sprint.Service
}

func NewSprintController(service sprint.Service) *SprintController {
	return &SprintController{
		sprintService: service,
	}
}

func (s *SprintController) Register(router *gin.RouterGroup) {
	router.POST("sprints", s.handleCreateSprint)
	router.GET("sprints", s.handleListSprint)
	router.GET("sprints/:sprintId", s.handleGetSprint)
	router.PATCH("sprints/:sprintId", s.handlePatchSprint)
}
