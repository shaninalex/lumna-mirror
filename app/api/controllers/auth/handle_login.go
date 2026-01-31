package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *AuthController) handleLogin(c *gin.Context) {
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

	accessTtl := time.Minute * 1
	token, err := auth.GenerateAccessJWTToken(identity.ID.String(), "all", accessTtl)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	toClient, toDB, err := auth.GenerateRefreshToken()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	refreshTtl := time.Hour * 24 * 30 // 30 days
	refreshExp := time.Now().Add(refreshTtl)
	rt := models.RefreshToken{
		IdentityID: identity.ID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	if result := db.GetDB(c.Request.Context()).Where("identity_id = ?", identity.ID.String()).Delete(&models.RefreshToken{}); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}
	if result := db.GetDB(c.Request.Context()).Create(&rt); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}

	c.SetCookie("access_token", token, int(accessTtl.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", toClient, int(refreshTtl.Seconds()), "/", "", true, true)

	utils.Success(c, identity, "Login successfull")
}
