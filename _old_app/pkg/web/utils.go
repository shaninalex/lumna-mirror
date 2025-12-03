package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/apperrors"
)

// ReturnJSON writes JSON response
func ReturnJSON(w http.ResponseWriter, status int, data any, params ...any) {
	resp := NewAPIResponse(data)
	if status >= 400 {
		resp.Status = false
	}

	for _, p := range params {
		switch v := p.(type) {
		case string:
			resp.Messages = append(resp.Messages, v)
		case apperrors.AppError:
			resp.Errors = append(resp.Errors, v)
		case error:
			var appErr apperrors.AppError
			if errors.As(v, &appErr) {
				resp.Errors = append(resp.Errors, appErr)
			} else {
				resp.Messages = append(resp.Messages, v.Error())
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
