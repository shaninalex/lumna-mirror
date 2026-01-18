package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/web/utils"

	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			utils.Error(c, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, internal.ContextUserID, uuid.MustParse(userID.(string)))
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
