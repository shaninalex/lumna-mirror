package config

type Database struct {
	SQlite   *DatabaseSqlite   `yaml:"sqlite,omitempty"`
	Postgres *DatabasePostgres `yaml:"postgres,omitempty"`
}

type DatabaseSqlite struct {
	Url string `yaml:"url"`
}

type DatabasePostgres struct {
	Host     string `host:"host"`
	Database string `host:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     uint16 `yaml:"port"`
}

type DatabaseType string

var (
	DatabaseTypeSQLite   DatabaseType = "sqlite"
	DatabaseTypePostgres DatabaseType = "postgres"
)

func DetectDatabase(c Database) DatabaseType {
	if c.SQlite != nil && c.Postgres == nil {
		return DatabaseTypeSQLite
	}

	if c.Postgres != nil && c.SQlite == nil {
		return DatabaseTypePostgres
	}

	panic("invalid database config")
}
