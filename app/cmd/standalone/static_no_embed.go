//go:build !embed
// +build !embed

package main

import "io/fs"

// GetStaticFS no embedded static files for test or CI builds
func getStaticFS() fs.FS {
	return nil
}
