package middlewares

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal"

	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, internal.ContextUserID, uuid.MustParse(userID.(string)))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
