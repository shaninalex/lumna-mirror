package utils

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderTemplate(ctx *gin.Context, httpStatus int, cmp templ.Component) {
	ctx.Status(httpStatus)
	templ.Handler(cmp).ServeHTTP(ctx.Writer, ctx.Request)
}
