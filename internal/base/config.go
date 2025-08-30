// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package base

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// NewConfig - new config.
func NewConfig(path string) *Config {
	conf := &Config{}
	conf.path = path
	conf.init()
	return conf
}

// IConfig - i config.
type IConfig interface {
	ReadConfig(path string)
	Env() string
	String(param string) string
	Bool(param string) bool
	Int(param string) int
	List(param string) []string
}

// Config - config.
type Config struct {
	path string
	v    *viper.Viper
}

func (s *Config) init() {
	s.v = viper.New()
	s.ReadConfig(s.path)
}

// ReadConfig - read config.
func (s *Config) ReadConfig(path string) {
	s.v.SetConfigFile(path)
	s.v.AutomaticEnv()
	s.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Handle errors reading the config file
	if err := s.v.ReadInConfig(); err != nil {
		// only fail if the error is not "file not found"
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("can't open config file: %w", err))
		}
	}
}

// NOTE: Not used for now
// const EnvProduction = "production"
// const EnvStaging = "staging"

const EnvDevelopment = "development"
const EnvTesting = "testing"

// Env - env.
func (s *Config) Env() string {
	if os.Getenv("APPLICATION_ENV") != "" {
		return os.Getenv("APPLICATION_ENV")
	}
	return EnvDevelopment
}

// String - string.
func (s *Config) String(param string) string {
	return s.v.GetString(param)
}

// Int - int.
func (s *Config) Int(param string) int {
	return s.v.GetInt(param)
}

// Bool - bool.
func (s *Config) Bool(param string) bool {
	return s.v.GetBool(param)
}

// List - lists all value.
func (s *Config) List(param string) []string {
	return s.v.GetStringSlice(param)
}
