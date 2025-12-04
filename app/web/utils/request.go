package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/shaninalex/lumna/app/base"
)

// GetUserID retrieves the user Id from context
func GetUserID(r *http.Request) uint {
	if id, ok := r.Context().Value(base.ContextUserID).(uint); ok {
		return id
	}
	panic("user was not found in request")
}

// GetAppName retrieves the app name from context
func GetAppName(r *http.Request) *string {
	if user, ok := r.Context().Value(base.ContextAppName).(*string); ok {
		return user
	}
	return nil
}

// BodyParser parse request POST body into generic type
func BodyParser[T any](r *http.Request) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func UrlNumericParam(w http.ResponseWriter, r *http.Request, name string) int64 {
	id, err := strconv.ParseInt(r.PathValue(name), 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return 0
	}
	return id
}

func UrlNumericQueryParam(r *http.Request, name string) int64 {
	id, err := strconv.ParseInt(r.URL.Query().Get(name), 10, 64)
	if err != nil {
		return 0
	}
	return id
}
