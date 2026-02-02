package auth

import (
	"errors"
	"log"
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
	refreshCookie, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println("[handleRefresh]: get refresh_token cookie error - ", err)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthRefreshToken)
		return
	}

	if refreshCookie == "" {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthRefreshToken)
		return
	}

	// find existed refresh token
	var dbRefreshToken *models.RefreshToken
	if result := db.GetDB(c.Request.Context()).
		Preload("Identity").
		Where("hash = ?", auth.ToHashToken(refreshCookie)).
		First(&dbRefreshToken); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	if dbRefreshToken == nil {
		utils.Error(c, http.StatusBadRequest, ErrorAuthRefreshToken)
		return
	}

	// Create new set of access/refresh tokens
	accessTtl := time.Minute * 15
	token, err := auth.GenerateAccessJWTToken(dbRefreshToken.IdentityID.String(), "all", accessTtl)
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
		IdentityID: dbRefreshToken.IdentityID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	// Delete existed refresh tokens
	if result := db.GetDB(c.Request.Context()).
		Where("identity_id = ?", dbRefreshToken.IdentityID.String()).
		Delete(&dbRefreshToken); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// create new one
	if result := db.GetDB(c.Request.Context()).Create(&rt); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}

	// assign new cookies
	c.SetCookie("access_token", token, int(accessTtl.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", toClient, int(refreshTtl.Seconds()), "/", "", true, true)

	utils.Success(c, nil)
}
