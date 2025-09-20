// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"net/http"

	"github.com/gorilla/csrf"
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers/templates"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func HandleLoginTemplate(w http.ResponseWriter, r *http.Request) {
	csrfInput := csrf.TemplateField(r)
	templates.LoginTemplate(string(csrfInput)).Render(r.Context(), w)
}

type loginForm struct {
	Email     string `form:"email"`
	Password  string `form:"password"`
	CsrfToken string `form:"csrf_token"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		web.Error(w, http.StatusBadRequest, err)
		return
	}

	formData := loginForm{
		Email:     r.PostFormValue("email"),
		Password:  r.PostFormValue("password"),
		CsrfToken: r.PostFormValue("csrf_token"),
	}
	web.Success(w, formData)
}
