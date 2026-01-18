package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:resources
var ResourcesFS embed.FS

// GetStaticFS embedded static files
func GetStaticFS() fs.FS {
	sub, err := fs.Sub(ResourcesFS, "resources/frontend/browser")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
