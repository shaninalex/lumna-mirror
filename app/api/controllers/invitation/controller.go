package invitation

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services/invitation"
	"gitlab.com/shaninalex/lumna/app/services/logger"
)

type InvitationController struct {
	invitationService invitation.Service
	logger            logger.Logger
}

func NewInvitationController(
	invitationService invitation.Service,
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
