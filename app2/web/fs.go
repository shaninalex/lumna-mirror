package web

import "embed"

//go:embed all:embed
var EmbedWebStaticFiles embed.FS
