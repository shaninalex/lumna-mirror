//go:build !embed

package lumna

import (
	"io/fs"
	"os"
)

// StaticFS returns static files read from disk (default build).
// Build with -tags embed to embed them into the binary instead.
func StaticFS(p string) fs.FS {
	return os.DirFS(p)
}
