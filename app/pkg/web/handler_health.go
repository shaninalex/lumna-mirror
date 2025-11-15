package web

import (
	"encoding/json"
	"net/http"

	"gitlab.com/shaninalex/lumna"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"name":    GetAppName(r),
		"version": lumna.Version,
	})
}
