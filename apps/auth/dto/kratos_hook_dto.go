// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"gitlab.com/shaninalex/flowreon/models"
)

// HooksKratosPayloadDTO - hooks kratos payload dto.
type HooksKratosPayloadDTO struct {
	UserID string            `json:"userId"`
	Traits models.UserTraits `json:"traits"`
}
