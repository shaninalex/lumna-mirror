package pages

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/root"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type PageController struct {
	localProvider *local.LocalAuthProvider
	userService   *services.UserService
}

func NewPageRouter(router *gin.RouterGroup, conf *config.Config) {
	controller := &PageController{
		localProvider: local.NewLocalAuthProvider(),
		userService:   &services.UserService{},
	}
	staticFS := lumna.StaticFS()
	if staticFS != nil {
		router.StaticFS("/assets", http.FS(staticFS))
		router.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(staticFS))
	} else {
		router.Static("/assets", "assets")                      // TODO: use assets base path from config
		router.StaticFile("/favicon.ico", "assets/favicon.ico") // TODO: use assets base path from config
	}
	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))
	router.Use(middlewares.CsrfMiddleware(conf.SecretKey))

	RegisterAuthPages(router)

	router.Use(middlewares.AuthRequired())
	router.Use(middlewares.IdentityMiddleware(controller.userService))
	{
		router.GET("/", controller.handleIndex)
		RegisterProjectPages(router)
		RegisterSettingsPages(router)
	}
}

func (s *PageController) handleIndex(ctx *gin.Context) {
	page := root.HomePageData{
		BasePage: utils.BasePageData(ctx, "Home page"),
	}
	utils.RenderTemplate(ctx, http.StatusOK, root.Home(page))
}
