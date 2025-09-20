// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"net/http"

	"github.com/gorilla/csrf"
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers/templates"
)

func (s *AuthHandler) HandleVerifyTemplate(w http.ResponseWriter, r *http.Request) {
	templates.VerifyTemplate(string(csrf.TemplateField(r)), nil).Render(r.Context(), w)
}

func (s *AuthHandler) HandleVerify(w http.ResponseWriter, r *http.Request) {
	templates.VerifyTemplate(string(csrf.TemplateField(r)), nil).Render(r.Context(), w)
}
