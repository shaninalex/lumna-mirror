// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers/templates"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthHandler) HandleLoginTemplate(w http.ResponseWriter, r *http.Request) {
	csrfInput := csrf.TemplateField(r)
	templates.LoginTemplate(string(csrfInput), nil).Render(r.Context(), w)
}

func (s *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := database.GetDb(r.Context())
	err := r.ParseForm()
	if err != nil {
		csrfInput := csrf.TemplateField(r)
		templates.LoginTemplate(string(csrfInput), err).Render(r.Context(), w)
		return
	}

	user, err := s.userRepository.GetByField(ctx, db, "email", r.PostFormValue("email"))
	if err != nil {
		templates.LoginTemplate(string(csrf.TemplateField(r)), err).Render(r.Context(), w)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(r.PostFormValue("password"))); err != nil {
		templates.LoginTemplate(string(csrf.TemplateField(r)), err).Render(r.Context(), w)
		return
	}

	session, _ := s.sessionStore.Get(r, "app_session")
	session.Values["user_id"] = user.ID
	session.Values["user_email"] = user.Email
	session.Options.MaxAge = 86400 * 7
	err = s.sessionStore.Save(r, w, session)
	if err != nil {
		csrfInput := csrf.TemplateField(r)
		templates.LoginTemplate(string(csrfInput), err).Render(r.Context(), w)
		return
	}

	log.Printf("user found: %s. Create and set session.\nRedirect to home page", user.Email)
	http.Redirect(w, r, "/", http.StatusFound)
}
