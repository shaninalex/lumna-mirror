package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
)

func (s *AuthController) HandleAuthLogin(c *gin.Context) {
	// get post payload
	payload := local.PasswordCredentials{}
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}

	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}

	// call authenticate
	authResult, err := s.localProvider.Authenticate(c.Request.Context(), &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", authResult.Identity.ID.String())
	session.Set("provider", authResult.Provider)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// return identity
	c.JSON(http.StatusOK, authResult.Identity)
}
