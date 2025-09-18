// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
)

// APIResponse standard response structure
type APIResponse struct {
	Status   bool                 `json:"status"`
	Data     any                  `json:"data,omitempty"`
	Messages []string             `json:"messages,omitempty"`
	Errors   []apperrors.AppError `json:"errors,omitempty"`
}
