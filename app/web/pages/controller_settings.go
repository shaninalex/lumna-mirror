package pages

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/settings"
)

type SettingsController struct{}

func NewSettingsController() *SettingsController {
	return &SettingsController{}
}

func ApplySettingsRoutes(router *gin.RouterGroup) {
	ctrl := NewSettingsController()

	router.GET("/settings", ctrl.handleIndex)
	router.GET("/settings/account", ctrl.handleAccount)
	router.GET("/settings/account/security", ctrl.handleSecurity)
}

func (c *SettingsController) handleIndex(ctx *gin.Context) {
	component := settings.Index()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *SettingsController) handleAccount(ctx *gin.Context) {
	component := settings.Account()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (c *SettingsController) handleSecurity(ctx *gin.Context) {
	component := templates.InitTemplate(ctx.Request.Context(), settings.Security)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}
