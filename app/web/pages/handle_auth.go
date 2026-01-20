package pages

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/pkg/tforms"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func (s *PageController) handleLogin(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	form := newLoginForm(utils.GetToken(c))
	component := templates.LoginPage(form)
	c.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(c.Writer, c.Request)
}

func (s *PageController) handleLoginSubmission(c *gin.Context) {
	payload := local.PasswordCredentials{}

	// Bind form data to struct
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := payload.Validate(); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// call authenticate
	authResult, err := s.localProvider.Authenticate(c.Request.Context(), &payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", authResult.Identity.ID.String())
	session.Set("provider", authResult.Provider)

	if err := session.Save(); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func newLoginForm(csrfToken string) tforms.IForm {
	form := tforms.NewForm(
		"/auth/login",
		false,
		tforms.NewHiddenField("_csrf", csrfToken, true),
		tforms.NewInputField("email", tforms.TextInputEmail, true).
			SetPlaceholder("example@mail.com").
			SetLabel("Email"),
		tforms.NewInputField("password", tforms.TextInputPassword, true).
			SetPlaceholder("Password").
			SetLabel("Password"),
	)

	return form
}
