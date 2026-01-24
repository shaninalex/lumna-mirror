//go:build !embed

package lumna

import (
	"io/fs"
)

// StaticFS returns nil when built without -tags embed
func StaticFS() fs.FS {
	return nil
}
