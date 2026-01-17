package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gitlab.com/shaninalex/lumna/app/internal/config"
)

func NewCookieStore(conf *config.Config) cookie.Store {
	store := cookie.NewStore([]byte(conf.SecretKey))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // TODO: from comfig ( 7 days  )
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // enable in HTTPS TODO: from comfig
	})
	return store
}
