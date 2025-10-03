// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"
)

type BadgeHandler struct{}

func NewBadgeHandler() *BadgeHandler {
	return &BadgeHandler{}
}

func (s *BadgeHandler) Post(w http.ResponseWriter, r *http.Request)   {}
func (s *BadgeHandler) Delete(w http.ResponseWriter, r *http.Request) {}
