package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

/*

Authorization: WRP+0Fs&@&KXiNS8>gV4_(;>4w]Ce*[ClJsa6MXXG+^t.2M[u+.7JD$FPd9[Jn#;J*5.plAy{z|Um{UV7N;bTb;S%l45Sz^2Jt#raxF9060YIFeYlKpz
Content-Type: application/json
Ory-Webhook-Request-Id: 5b1c5716-db2a-42b6-bc77-7b7619f73d86

*/

// AuthHooksMiddleware - auth middleware.
type AuthHooksMiddleware struct {
	kratosService kratos.IKratos
	config        base.IConfig
}

// NewAuthHooksMiddleware - new auth hooks middleware.
func NewAuthHooksMiddleware(kratosService kratos.IKratos, config base.IConfig) fiber.Handler {
	m := &AuthHooksMiddleware{
		kratosService: kratosService,
		config:        config,
	}
	return m.Wrap()
}

// Wrap - wrapper, actual middleware
func (s *AuthHooksMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Get("Ory-Webhook-Request-Id")
		if requestID == "" {
			return web.Error(ctx, http.StatusBadRequest, fmt.Errorf("request id no set"))
		}
		if ctx.Get("Authorization") != s.config.String("app.secret_key") {
			return web.Error(ctx, http.StatusUnauthorized, fmt.Errorf("unauthorized request"))
		}
		log.Println("AuthHooksMiddleware handle request:", ctx.Get("Host"))
		return ctx.Next()
	}
}
