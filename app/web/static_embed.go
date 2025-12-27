//go:build embed
// +build embed

package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:embed/browser
var webFS embed.FS

// GetStaticFS embedded static files
func GetStaticFS() fs.FS {
	sub, err := fs.Sub(webFS, "embed/frontend/browser")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
