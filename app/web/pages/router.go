package pages

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/static"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates"
)

type PageController struct {
}

func NewPageRouter(router *gin.RouterGroup) {
	controller := &PageController{}
	staticFiles, err := fs.Sub(static.GetStaticFS(), "assets")
	if err != nil {
		panic(err)
	}
	router.StaticFS("/assets", http.FS(staticFiles))
	router.GET("/", controller.handleIndex)
	router.GET("/auth/login", controller.handleLogin)
}

func (s *PageController) handleIndex(ctx *gin.Context) {
	component := templates.Index()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (s *PageController) handleLogin(ctx *gin.Context) {
	component := templates.LoginPage()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}
