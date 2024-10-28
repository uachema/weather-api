package util

import (
	"encoding/json"
	"net/http"
)

// writeJSON encodes a given value as JSON and writes it to the http.ResponseWriter with the specified HTTP status code.
// It sets the "Content-Type" header to "application/json" and returns any encoding error encountered.
//
// Parameters:
// - w: The http.ResponseWriter for writing the response.
// - status: The HTTP status code to send.
// - v: The value to encode as JSON.
//
// Returns:
// - error: An error if JSON encoding fails.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
