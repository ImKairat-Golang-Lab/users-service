package http

import (
	"net/http"
	"encoding/json"
)


func logFields(r *http.Request, component string, statusCode int) map[string]any {
	return map[string]any{
		"component":   component,
		"status_code": statusCode,
		"method":      r.Method,
		"endpoint":    r.URL.Path,
		"user_agent":  r.UserAgent(),
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}