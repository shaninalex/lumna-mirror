// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package dto

import (
	"time"
)

type OrganizationDto struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
