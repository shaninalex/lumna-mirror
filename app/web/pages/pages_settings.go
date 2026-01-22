package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/settings"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type SettingsController struct {
}

func NewSettingsController() *SettingsController {
	return &SettingsController{}
}
func RegisterSettingsPages(router *gin.RouterGroup) {
	controller := NewSettingsController()

	router.GET("/settings", controller.settingsIndex)
}

func (s *SettingsController) settingsIndex(c *gin.Context) {
	base := utils.GetBasePage(c.Request.Context())
	base.Title = "Settings"
	pageData := settings.SettingsPageData{
		BasePage: base,
	}
	utils.RenderTemplate(c, http.StatusOK, settings.Index(pageData))
}
