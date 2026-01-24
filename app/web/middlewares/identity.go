package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/services"
)

func IdentityMiddleware(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		identity, _ := s.Identity(c.Request.Context())
		if identity != nil {
			ctx := context.WithValue(
				c.Request.Context(),
				internal.ContextIdentity,
				identity,
			)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
