package inventory

import (
	"encoding/json"
	"net/http"
)

// handleHealthCheck is a simple handler to check if the service is running.
func (s *Server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a response payload
		response := map[string]string{
			"status": "ok",
		}

		// Set the content-type header and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
