//go:build embed
// +build embed

package main

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:web/browser
var webFS embed.FS

// GetStaticFS embedded static files
func getStaticFS() fs.FS {
	sub, err := fs.Sub(webFS, "web/browser")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
