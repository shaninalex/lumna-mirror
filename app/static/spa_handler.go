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

func SPAHandler(staticFS fs.FS) gin.HandlerFunc {
	httpFS := http.FS(staticFS)

	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		filePath := strings.TrimPrefix(reqPath, "/")

		if filePath == "" {
			filePath = "index.html"
		}

		// Try to serve the exact file if it exists
		if f, err := staticFS.Open(filePath); err == nil {
			f.Close()
			// Set Content-Type explicitly for JS files (ES modules require correct MIME)
			if ext := path.Ext(filePath); ext == ".js" || ext == ".mjs" {
				c.Header("Content-Type", "application/javascript")
			}
			c.FileFromFS(reqPath, httpFS)
			return
		}

		// SPA fallback: serve index.html for client-side routing
		c.FileFromFS("index.html", httpFS)
	}
}
