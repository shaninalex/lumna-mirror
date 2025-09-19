// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"encoding/json"
	"net/http"

	"gitlab.com/shaninalex/flowreon"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"name":    GetAppName(r),
		"version": flowreon.Version,
	})
}
