// Create default config file if not exists

package base

import (
	"log"
	"os"

	"gitlab.com/shaninalex/lumna/app/internal/dir"
	"gitlab.com/shaninalex/lumna/app/internal/utils"
	"gopkg.in/yaml.v2"
)

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
