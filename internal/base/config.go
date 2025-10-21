// Copyright © 2025 Lumna. All rights reserved.

package base

import (
	"github.com/shaninalex/lumna/internal/utils"
)

// IsDebug - return true if environment is in development mode
func IsDebug() bool {
	env := utils.GetEnv("LUMNA_ENVIRONMENT", "development")
	return env == "development"
}
