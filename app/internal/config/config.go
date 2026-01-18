package config

import (
	"context"
	"io"
	"os"

	"gitlab.com/shaninalex/lumna/app/internal"
	"gopkg.in/yaml.v3"
)

type Environment string

const (
	EnvironmentDev  Environment = "dev"
	EnvironmentTest Environment = "test"
)

type Serve struct {
	Port int `yaml:"port"`
}

type Database struct {
	Url string `yaml:"url"`
}

type Config struct {
	Env       Environment `yaml:"env"`
	Serve     Serve       `yaml:"serve"`
	Database  Database    `yaml:"database"`
	SecretKey string      `yaml:"secret_key"`
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

func GetConfig(ctx context.Context) *Config {
	cnf, ok := ctx.Value(internal.ContextConfig).(*Config)
	if !ok {
		panic("config not found in context")
	}
	return cnf
}
