package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

var (
	ErrorAuthRefreshToken = errors.New("invalid refresh token")
)

func (s *AuthController) handleRefresh(c *gin.Context) {
	userID, _ := utils.GetUserID(c)
	refrashCookie, err := c.Cookie("refresh_token")
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// find and delete existed refresh token
	var dbRefreshToken *models.RefreshToken
	if result := db.GetDB(c.Request.Context()).Where("hash = ?", auth.ToHashToken(refrashCookie)).First(&dbRefreshToken); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	if dbRefreshToken == nil {
		utils.Error(c, http.StatusBadRequest, ErrorAuthRefreshToken)
		return
	}

	if result := db.GetDB(c.Request.Context()).Where("id = ?", dbRefreshToken.ID.String()).Delete(&dbRefreshToken); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// Create new set of access/refresh tokens
	accessTtl := time.Minute * 15
	token, err := auth.GenerateAccessJWTToken(userID.String(), "all", accessTtl)
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
		IdentityID: userID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	if result := db.GetDB(c.Request.Context()).Create(&rt); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}

	// assign new cookies
	c.SetCookie("access_token", token, int(accessTtl.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", toClient, int(refreshTtl.Seconds()), "/", "", true, true)

	utils.Success(c, nil)
}
