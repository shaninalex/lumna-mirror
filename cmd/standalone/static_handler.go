// Copyright © 2025 Lumna. All rights reserved.

package main

import (
	"io/fs"
	"net/http"
	"strings"
)

func frontendHandler(static fs.FS) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "index.html"
		}
		_, err := static.Open(strings.TrimPrefix(path, "/"))
		if err == nil {
			http.FileServer(http.FS(static)).ServeHTTP(w, r)
			return
		}
		http.ServeFileFS(w, r, static, "index.html")
	}
}
