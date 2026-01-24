package pages

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/auth"
	"gitlab.com/shaninalex/lumna/app/web/pages/templates/partials"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type AuthController struct {
	localProvider *local.LocalAuthProvider
}

func NewAuthController() *AuthController {
	return &AuthController{
		localProvider: local.NewLocalAuthProvider(),
	}
}

func RegisterAuthPages(router *gin.RouterGroup) {
	auth := NewAuthController()
	router.GET("/auth/login", auth.handleLogin)
	router.POST("/auth/login", auth.handleLoginSubmission)
	router.GET("/auth/logout", auth.handleLogoutSubmission)
}

func (s *AuthController) handleLogin(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	page := utils.BasePageData(c, "Login")
	component := auth.LoginPage(page)
	templ.Handler(component).ServeHTTP(c.Writer, c.Request)
}

func (s *AuthController) handleLogoutSubmission(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	_ = session.Save()
	c.Redirect(http.StatusFound, "/auth/login")
}

func (s *AuthController) handleLoginSubmission(c *gin.Context) {
	payload := local.PasswordCredentials{}

	// Bind form data to struct
	if err := c.ShouldBind(&payload); err != nil {
		utils.RenderTemplate(c, http.StatusOK, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	if err := payload.Validate(); err != nil {
		utils.RenderTemplate(c, http.StatusOK, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	// call authenticate
	authResult, err := s.localProvider.Authenticate(c.Request.Context(), &payload)
	if err != nil {
		utils.RenderTemplate(c, http.StatusOK, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", authResult.Identity.ID.String())
	session.Set("provider", authResult.Provider)

	if err := session.Save(); err != nil {
		utils.RenderTemplate(c, http.StatusOK, partials.Alert(err.Error(), &partials.AlertTypeDanger))
		return
	}

	// redirect to home page
	c.Writer.Header().Set("HX-Redirect", "/")
}
