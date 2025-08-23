// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package tdata

import "gitlab.com/shaninalex/jajirra/internal/base"

func newTestConfig() base.IConfig {
	conf = &testConfig{
		storage: map[string]any{
			"app.dsn": "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=test",
		},
	}
	return conf
}

var conf *testConfig

type testConfig struct {
	storage map[string]any
}

func (s *testConfig) ReadConfig(path string) {
	// TODO: read from test config file
}

func (s *testConfig) Env() string { return base.ENV_TESTING }

func (s *testConfig) String(param string) string {
	v, ok := s.storage[param].(string)
	if !ok {
		return ""
	}
	return v
}

func (s *testConfig) Bool(param string) bool {
	v, ok := s.storage[param].(bool)
	if !ok {
		return false
	}
	return v
}

func (s *testConfig) Int(param string) int {
	v, ok := s.storage[param].(int)
	if !ok {
		return 0
	}
	return v
}

func (s *testConfig) List(param string) []string {
	v, ok := s.storage[param].([]string)
	if !ok {
		return []string{}
	}
	return v
}

func (s *testConfig) Set(k string, v any) {
	s.storage[k] = v
}

func SetConfigValue(k string, v any) {
	if conf == nil {
		conf = &testConfig{
			storage: make(map[string]any),
		}
	}
	conf.Set(k, v)
}
