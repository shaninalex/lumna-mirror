//go:build embed

package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:resources
var resourcesFS embed.FS

// GetStaticFS returns embedded static files when built with -tags embed
func GetStaticFS() fs.FS {
	sub, err := fs.Sub(resourcesFS, "resources/frontend/browser")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
