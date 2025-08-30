// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package utils

import (
	"fmt"
	"math/rand/v2"
)

func GenerateEntityCode(entity string) string {
	n := rand.IntN(1_000_000)
	return fmt.Sprintf("%s-%06d", entity, n)
}
