package utils

import (
	"fmt"
	"math/rand/v2"
)

// GenerateEntityCode - generate entity code.
func GenerateEntityCode(entity string) string {
	n := rand.IntN(1_000_000)
	return fmt.Sprintf("%s-%06d", entity, n)
}
