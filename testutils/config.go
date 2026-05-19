package testutils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gitlab.com/shaninalex/lumna/app/pkg/config"
)

func testConfigPath() string {
	if os.Getenv("LUMNA_TEST_CONFIG") != "" {
		if _, err := os.Stat(os.Getenv("LUMNA_TEST_CONFIG")); err != nil {
			if os.IsNotExist(err) {
				panic(fmt.Errorf("path %s is not exists", os.Getenv("LUMNA_TEST_CONFIG")))
			}
		}
		return os.Getenv("LUMNA_TEST_CONFIG")
	}

	_, thisFile, _, _ := runtime.Caller(0)
	repoRoot := filepath.Join(filepath.Dir(thisFile), "..")
	return filepath.Join(repoRoot, "config", "config.test.yaml")
}

func ProvideTestConfig() *config.Config {
	return config.ReadConfig(testConfigPath())
}
