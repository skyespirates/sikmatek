package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]any{
		"message": message,
		"data":    data,
	}
	_ = json.NewEncoder(w).Encode(resp)
}
