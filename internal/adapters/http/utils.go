package http

import (
	"encoding/json"
	"net/http"

	ports "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
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

func writeJSON(w http.ResponseWriter, logger ports.Logger, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Warn("failed to encode JSON response", map[string]any{
			"error": err.Error(),
		})
	}
}
