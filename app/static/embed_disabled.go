//go:build !embed

package static

import (
	"embed"
	"io/fs"
	"log"
)

// // GetStaticFS returns nil when built without -tags embed
// func GetStaticFS() fs.FS {
// 	return nil
// }

//go:embed all:resources
var resourcesFS embed.FS

// GetStaticFS returns embedded static files when built with -tags embed
func GetStaticFS() fs.FS {
	sub, err := fs.Sub(resourcesFS, "resources")
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}
	return sub
}
