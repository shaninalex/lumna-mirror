//go:build embed

package lumna

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:resources/assets
//go:embed all:resources/migrations
var resourcesFS embed.FS

// StaticFS returns embedded static files when built with -tags embed
func StaticFS(p string) fs.FS {
	sub, err := fs.Sub(resourcesFS, p)
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}

	return sub
}
