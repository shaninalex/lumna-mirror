package utils

import (
	"os"
)

// GetEnv - get env variable with default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
