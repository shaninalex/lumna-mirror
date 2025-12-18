//go:build !embed
// +build !embed

package web

import "io/fs"

// GetStaticFS no embedded static files for test or CI builds
func GetStaticFS() fs.FS {
	return nil
}
