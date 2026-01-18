package web

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Disable automatic redirects to allow SPA fallback routing
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	return router
}
