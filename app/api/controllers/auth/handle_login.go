package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/jwt"
)

func (s *AuthController) HandleAuthLogin(c *gin.Context) {
	payload := local.PasswordCredentials{}
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	identity, err := s.localProvider.Authenticate(c.Request.Context(), &payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	token, err := jwt.GenerateLoginToken(identity.ID.String())
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, map[string]string{
		"login_token": token,
	})
}
