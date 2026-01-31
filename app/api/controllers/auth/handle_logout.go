package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *AuthController) handleLogout(c *gin.Context) {
	userID, _ := utils.GetUserID(c)
	if result := db.GetDB(c.Request.Context()).Where("identity_id = ?", userID.String()).Delete(&models.RefreshToken{}); result.Error != nil {
		utils.Error(c, http.StatusBadRequest, result.Error)
		return
	}

	c.SetCookie("access_token", "", 0, "/", "", true, true)
	c.SetCookie("refresh_token", "", 0, "/", "", true, true)

	utils.Success(c, nil, "Logout successfull")
}
