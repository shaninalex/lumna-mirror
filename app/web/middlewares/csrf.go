package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func CsrfMiddleware(secretKey string) gin.HandlerFunc {
	return utils.CSRFMiddleware(utils.Options{
		Secret: secretKey,
		ErrorFunc: func(c *gin.Context) {
			utils.Error(c, http.StatusForbidden, fmt.Errorf("CSRF token mismatch"))
			c.Abort()
		},
	})
}
