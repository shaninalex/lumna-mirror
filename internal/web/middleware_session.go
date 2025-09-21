package web

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

func SessionMiddleware(store *CookieStoreDatabase, sessionName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, sessionName)
			if err != nil {
				http.Error(w, "Session error", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), base.ContextSession, session)
			next.ServeHTTP(w, r.WithContext(ctx))
			if session.IsNew || len(session.Values) > 0 {
				if err := store.Save(r, w, session); err != nil {
					fmt.Println("Failed to save session:", err)
				}
			}
		})
	}
}

func AuthSessionMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session := GetSession(r)
			if session == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, ok := session.Values["user_id"].(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			_userID, err := uuid.Parse(userID)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			user, err := repositories.NewUserRepository().GetByField(r.Context(), db, "id", _userID.String())
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), base.ContextUser, user)
			ctx = context.WithValue(ctx, base.ContextUserID, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
