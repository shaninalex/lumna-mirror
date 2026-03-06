package invitation

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gitlab.com/shaninalex/lumna/app/services"
)

type InvitationController struct {
	invitationService *services.InvitationService
	logger            logger.Logger
}

func NewInvitationController(
	invitationService *services.InvitationService,
	logger logger.Logger,
) *InvitationController {
	s := &InvitationController{
		invitationService: invitationService,
		logger:            logger,
	}

	return s
}

func (s *InvitationController) Register(router *gin.RouterGroup) {
	router.GET("invitations", s.List)
	router.POST("invitations", s.Create)
}
