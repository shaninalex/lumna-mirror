package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:resources
var resourcesFS embed.FS

func GetStaticFS() fs.FS {
	sub, err := fs.Sub(resourcesFS, "resources")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
