package api

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/base"
)

// GetAppName retrieves the app name from context
func GetAppName(r *http.Request) *string {
	if user, ok := r.Context().Value(base.ContextAppName).(*string); ok {
		return user
	}
	return nil
}
