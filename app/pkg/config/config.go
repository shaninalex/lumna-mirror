package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Environment string

const (
	EnvironmentDev  Environment = "dev"
	EnvironmentTest Environment = "test"
)

type Serve struct {
	Port  int  `yaml:"port"`
	Embed bool `yaml:"embed"`
}

type Database struct {
	Url string `yaml:"url"`
}

type EmailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Env       Environment `yaml:"env"`
	Serve     Serve       `yaml:"serve"`
	Database  Database    `yaml:"database"`
	SecretKey string      `yaml:"secret_key"`
	Email     EmailConfig `yaml:"email"`
}

func ReadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	return &config
}

func ProvideConfig(configPath string) func() *Config {
	return func() *Config {
		return ReadConfig(configPath)
	}
}
