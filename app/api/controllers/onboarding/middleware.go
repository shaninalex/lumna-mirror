package onboarding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *OnboardingController) workspaceExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		wp, err := s.workspaceManager.List(c.Request.Context())
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, err)
			return
		}
		if len(wp) > 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func (s *OnboardingController) teammatesExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.userManager.List(c.Request.Context())
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, err)
			return
		}
		if len(users) > 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		invs, err := s.invitationManager.List(c.Request.Context())
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, err)
			return
		}
		if len(invs) > 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
