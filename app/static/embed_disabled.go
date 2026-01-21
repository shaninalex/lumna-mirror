//go:build !embed

package static

import (
	"io/fs"
)

// GetStaticFS returns nil when built without -tags embed
func GetStaticFS() fs.FS {
	return nil
}
