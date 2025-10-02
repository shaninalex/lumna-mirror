// Copyright © 2025 Lumna. All rights reserved.

package models

import (
	"time"
)

type Project struct {
	ID        uint
	Title     string
	Code      string
	Statuses  []Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID    uint
	Title string
	Idx   uint
}
