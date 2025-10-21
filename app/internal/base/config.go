// Copyright © 2025 Lumna. All rights reserved.

package base

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/shaninalex/lumna/app/internal/dir"
	"github.com/spf13/viper"
)

const EnvProduction = "production"
const EnvDevelopment = "development"
const EnvTesting = "testing"

var config *Config

func GetConfig() *Config {
	return config
}

func init() {
	fmt.Print("Init config ... ")

	config = NewConfig()

	fmt.Print("ok\n")
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
	s.startTime = time.Now()
	s.v = viper.New()

	configPath := dir.ConfigDirectory
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		configPath = os.Getenv("LUMNA_CONFIG_PATH")
	}
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		panic(err)
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
