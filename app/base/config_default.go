package base

import (
	"crypto/rand"
	"log"
	"os"

	"gitlab.com/shaninalex/lumna/app/pkg/dir"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gopkg.in/yaml.v2"
)

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

func CreateDefaultConfig(configFilePath string) error {
	key, err := generaterandomstring(64)
	if err != nil {
		return err
	}
	conf := &ConfigModel{
		DatabasePath: dir.DefaultDatabasePath(),
		SecretKey:    key,
		Port:         utils.Pointer[int](8000),
	}

	file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(conf)
}

func generaterandomstring(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}
