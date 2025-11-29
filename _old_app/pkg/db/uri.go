package db

import (
	"path"
	"strings"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/base"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/dir"
)

type DatabaseProvider string

const (
	SQLite   DatabaseProvider = "sqlite"
	Postgres DatabaseProvider = "postgres"
)

// GetDatabaseUri parse database uri
func GetDatabaseUri(config *base.Config) string {
	var provider DatabaseProvider

	if strings.HasPrefix(config.String("database_path"), "postgres") {
		provider = Postgres
	} else {
		provider = SQLite
	}

	return buildPath(config, provider)
}

func buildPath(config *base.Config, provider DatabaseProvider) string {
	if provider == SQLite {
		return path.Join(dir.GetShareDir(), config.String("database_path"))
	}

	return config.String("database_path")
}
