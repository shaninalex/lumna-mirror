// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"
)

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (s *StatusHandler) Patch(w http.ResponseWriter, r *http.Request) {}
