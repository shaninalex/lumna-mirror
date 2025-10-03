// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"net/http"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (s *TaskHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (s *TaskHandler) Patch(w http.ResponseWriter, r *http.Request)  {}
func (s *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {}
