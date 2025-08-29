// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package base

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig(path string) *Config {
	conf := &Config{}
	conf.path = path
	conf.init()
	return conf
}

func ProvideConfig(path string) IConfig {
	return NewConfig(path)
}

type IConfig interface {
	ReadConfig(path string)
	Env() string
	String(param string) string
	Bool(param string) bool
	Int(param string) int
	List(param string) []string
}

type Config struct {
	path string
	v    *viper.Viper
}

func (s *Config) init() {
	s.v = viper.New()
	s.ReadConfig(s.path)
}

func (s *Config) ReadConfig(path string) {
	s.v.SetConfigFile(path)

	// Handle errors reading the config file
	if err := s.v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("can't open config file. %w", err))
	}
}

const ENV_PRODUCTION = "production"
const ENV_STAGING = "staging"
const ENV_DEVELOPMENT = "development"
const ENV_TESTING = "testing"

func (s *Config) Env() string {
	if os.Getenv("APPLICATION_ENV") != "" {
		return os.Getenv("APPLICATION_ENV")
	}
	return ENV_DEVELOPMENT
}

func (s *Config) String(param string) string {
	return s.v.GetString(param)
}

func (s *Config) Int(param string) int {
	return s.v.GetInt(param)
}

func (s *Config) Bool(param string) bool {
	return s.v.GetBool(param)
}

func (s *Config) List(param string) []string {
	return s.v.GetStringSlice(param)
}
