package base

type ConfigModel struct {
	// DatabasePath is a path for database sqlite3 file
	// or URI connection string for postgres / mysql
	DatabasePath string `yaml:"database_path"`

	// SecretKey used for web authentication
	SecretKey string `yaml:"secret_key"`

	// Mode working application mode such as "development", "production" ( default ) or "testing"
	Mode *string `yaml:"mode,omitempty"`

	// Port working server port number
	Port *int `yaml:"port,omitempty"`
}
