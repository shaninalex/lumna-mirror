package base

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/shaninalex/lumna/app/internal/dir"
)

const EnvProduction = "production"
const EnvDevelopment = "development"
const EnvTesting = "testing"

var config *Config

func GetConfig() *Config {
	return config
}

func init() {
	config = NewConfig()
}

func NewConfig() *Config {
	conf := &Config{}
	conf.init()
	return conf
}

type Config struct {
	v         *viper.Viper
	startTime time.Time
}

func (s *Config) init() {
	if err := dir.MakeProjectDirectories(); err != nil {
		panic(err)
	}

	s.startTime = time.Now()
	s.v = viper.New()

	configPath := os.Getenv("LUMNA_CONFIG_PATH")

	s.ReadConfig(configPath)
}

func (s *Config) ReadConfig(configPath string) {
	if _, err := os.Stat(configPath); err != nil {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configPath = path.Join(home, dir.ConfigDirectory, "config.yaml")
	}

	s.v.SetConfigFile(configPath)
	if err := s.v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Can't open config file %s. ERROR: %w \n", configPath, err))
	}
}
func (s *Config) Env() string {
	return s.v.GetString("mode")
}

func (s *Config) String(param string) string {
	return s.v.GetString(param)
}

func (s *Config) Int(param string) int {
	return s.v.GetInt(param)
}

func (s *Config) List(param string) []string {
	return s.v.GetStringSlice(param)
}

// IsDebug - return true if environment is in development mode
func IsDebug() bool {
	return config.Env() == "development"
}
