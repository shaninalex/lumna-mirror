package base

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	if env := os.Getenv("LUMNA_CONFIG_PATH"); env != "" {
		if fi, err := os.Stat(env); err == nil && fi.IsDir() {
			// if config file does not exists
			if _, err := os.Stat(path.Join(env, "config.yaml")); err != nil {
				CreateDefaultConfig(path.Join(env, "config.yaml"))
			}
			s.ReadConfig(env)
			return
		} else if err != nil {
			panic(fmt.Errorf("LUMNA_CONFIG_PATH invalid: %w", err))
		} else {
			panic(fmt.Errorf("LUMNA_CONFIG_PATH is not a directory: %s", env))
		}
	}

	// Use standard $HOME/.config/lumna path if env was not set.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(home, dir.ConfigDirectory)
	if fi, err := os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			panic(fmt.Errorf("standard config directory does not exist: %s", configPath))
		}
		panic(fmt.Errorf("cannot stat standard config directory %s: %w", configPath, err))
	} else if !fi.IsDir() {
		panic(fmt.Errorf("standard config path is not a directory: %s", configPath))
	}

	// if config file does not exists
	if _, err := os.Stat(path.Join(configPath, "config.yaml")); err != nil {
		CreateDefaultConfig(path.Join(configPath, "config.yaml"))
	}

	s.ReadConfig(configPath)
}

func (s *Config) ReadConfig(path string) {
	s.v.AddConfigPath(path)
	s.v.SetConfigType("yaml")
	s.v.SetConfigName("config")

	err := s.v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Can't open config file. %w \n", err))
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
