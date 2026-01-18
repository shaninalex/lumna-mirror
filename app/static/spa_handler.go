package static

import (
	"io/fs"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	// Ensure correct MIME types for web assets (required for ES modules)
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".mjs", "application/javascript")
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".json", "application/json")
	mime.AddExtensionType(".woff2", "font/woff2")
	mime.AddExtensionType(".woff", "font/woff")
	mime.AddExtensionType(".svg", "image/svg+xml")
}

const spaPrefix = "frontend/browser"

func SPAHandler(staticFS fs.FS) gin.HandlerFunc {
	// Pre-read index.html for SPA fallback (avoids http.FileServer redirect issues)
	indexHTML, err := fs.ReadFile(staticFS, spaPrefix+"/index.html")
	if err != nil {
		panic("failed to read index.html from embedded files: " + err.Error())
	}

	// Create a sub-filesystem for the SPA files
	spaFS, err := fs.Sub(staticFS, spaPrefix)
	if err != nil {
		panic("failed to create SPA sub-filesystem: " + err.Error())
	}
	httpFS := http.FS(spaFS)

	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		filePath := strings.TrimPrefix(reqPath, "/")

		if filePath == "" {
			filePath = "index.html"
		}

		// Try to serve the exact file if it exists and is not a directory
		if f, err := spaFS.Open(filePath); err == nil {
			stat, statErr := f.Stat()
			f.Close()
			if statErr == nil && !stat.IsDir() {
				// Set Content-Type explicitly for JS files (ES modules require correct MIME)
				if ext := path.Ext(filePath); ext == ".js" || ext == ".mjs" {
					c.Header("Content-Type", "application/javascript")
				}
				c.FileFromFS(reqPath, httpFS)
				return
			}
		}

		// SPA fallback: serve index.html directly for client-side routing
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	}
}
