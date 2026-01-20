package pages

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/static"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates"
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
	staticFiles, err := fs.Sub(static.GetStaticFS(), "assets")
	if err != nil {
		panic(err)
	}

	router.Use(sessions.Sessions("session", utils.NewCookieStore(conf)))
	router.Use(utils.CSRFMiddleware(utils.Options{
		Secret: conf.SecretKey,
		ErrorFunc: func(c *gin.Context) {
			utils.Error(c, http.StatusForbidden, fmt.Errorf("CSRF token mismatch"))
			c.Abort()
		},
	}))
	router.StaticFS("/assets", http.FS(staticFiles))

	router.GET("/auth/login", controller.handleLogin)
	router.POST("/auth/login", controller.handleLoginSubmission)
	router.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(staticFiles))

	router.Use(middlewares.AuthRequired())
	{
		router.GET("/", controller.handleIndex)
	}
}

func (s *PageController) handleIndex(ctx *gin.Context) {
	component := templates.Index()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}
