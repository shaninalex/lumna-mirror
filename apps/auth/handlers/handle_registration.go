// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/csrf"
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers/templates"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthHandler) HandleRegistrationTemplate(w http.ResponseWriter, r *http.Request) {
	templates.RegistrationTemplate(string(csrf.TemplateField(r)), nil).Render(r.Context(), w)
}

func (s *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := database.GetDb(r.Context())
	err := r.ParseForm()
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}
	_, err = s.userRepository.GetByField(ctx, db, "email", r.PostFormValue("email"))
	if err == nil {
		csrfInput := csrf.TemplateField(r)
		templates.RegistrationTemplate(string(csrfInput), errors.New("user already exists")).Render(r.Context(), w)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		templates.RegistrationTemplate(string(csrf.TemplateField(r)), err).Render(r.Context(), w)
		return
	}

	user := &models.User{
		Email:        r.PostFormValue("email"),
		PasswordHash: string(hash),
	}
	user, err = s.userRepository.Save(ctx, db, user)
	if err != nil {
		templates.RegistrationTemplate(string(csrf.TemplateField(r)), err).Render(r.Context(), w)
		return
	}

	http.Redirect(w, r, "/auth/login", http.StatusFound)
}
