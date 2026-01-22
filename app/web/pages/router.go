package pages

import (
	"io/fs"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/static"
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
	staticFilesFS := static.GetStaticFS()
	if staticFilesFS != nil {
		staticFiles, err := fs.Sub(staticFilesFS, "assets")
		if err != nil {
			panic(err)
		}
		router.StaticFS("/assets", http.FS(staticFiles))
		router.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(staticFiles))
	} else {
		router.Static("/assets", "app/static/resources/assets")
		router.StaticFile("/favicon.ico", "app/static/resources/assets/favicon.ico")
	}

	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))
	router.Use(middlewares.CsrfMiddleware(conf.SecretKey))

	RegisterAuthPages(router)

	router.Use(middlewares.AuthRequired())
	router.Use(middlewares.IdentityMiddleware(controller.userService))
	{
		router.GET("/", controller.handleIndex)
	}
}

func (s *PageController) handleIndex(ctx *gin.Context) {
	page := root.HomePageData{
		BasePage: utils.GetBasePage(ctx.Request.Context()),
	}
	utils.RenderTemplate(ctx, http.StatusOK, root.Home(page))
}
