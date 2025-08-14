// Copyright © 2025 Soundstream https://soundstream.shaninalex.com. All rights reserved.

package tdata

import "gitlab.com/shaninalex/jajirra/internal/base"

func newTestConfig() base.IConfig {
	return &testConfig{}
}

type testConfig struct {
}

func (s *testConfig) ReadConfig(path string)     {}
func (s *testConfig) Env() string                { return base.ENV_TESTING }
func (s *testConfig) String(param string) string { return "" }
func (s *testConfig) Bool(param string) bool     { return true }
func (s *testConfig) Int(param string) int       { return 0 }
func (s *testConfig) List(param string) []string { return []string{} }
func (s *testConfig) AudioStorage() string       { return "" }
func (s *testConfig) ImageStorage() string       { return "" }
