package api

import (
	"encoding/json"
	"net/http"

	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

func HealthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status":  "ok",
		"name":    utils.GetAppName(r),
		"version": lumna.Version,
	})
}
