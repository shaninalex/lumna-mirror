package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *AuthController) handleLogout(c *gin.Context) {
	userID, _ := utils.GetUserID(c)
	// TODO: move to auth service. Do not call db in handlers directly!
	if result := persistence.GetDB(c.Request.Context()).Where("identity_id = ?", userID.String()).Delete(&models.RefreshToken{}); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}

	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	utils.Success(c, nil, "Logout successfull")
}
