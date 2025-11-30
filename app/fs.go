package app

import "embed"

//go:embed all:migrations
var EmbedDatabaseMigrations embed.FS

//go:embed all:web/embed
var EmbedWebStaticFiles embed.FS
