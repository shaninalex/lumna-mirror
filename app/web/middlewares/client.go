package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ClientMiddleware copy context values from client to gin context
func ClientMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
