package pages

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/pkg/tforms"
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
	router.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(staticFiles))
}

func (s *PageController) handleIndex(ctx *gin.Context) {
	component := templates.Index()
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func (s *PageController) handleLogin(ctx *gin.Context) {
	form := newLoginForm()
	component := templates.LoginPage(form)
	ctx.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(ctx.Writer, ctx.Request)
}

func newLoginForm() tforms.IForm {
	form := tforms.NewForm(
		"/auth/login",
		false,
		tforms.NewHiddenField("_csrf", "csrf-token-string", true),
		tforms.NewInputField("email", tforms.TextInputEmail, true).
			SetPlaceholder("example@mail.com").
			SetLabel("Email"),
		tforms.NewInputField("password", tforms.TextInputPassword, true).
			SetPlaceholder("Password").
			SetLabel("Password"),
	)

	return form
}
