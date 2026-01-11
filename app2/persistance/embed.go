package persistance

import "embed"

//go:embed all:migrations
var EmbedDatabaseMigrations embed.FS
