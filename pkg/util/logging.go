package util

import (
	"log"
	"net/http"
)

// logRequest logs details about an incoming HTTP request, including method, path, and remote IP address.
//
// Parameters:
// - r: The http.Request to log.
func LogRequest(r *http.Request) {
	log.Printf("Request received - Method: %s, Path: %s, RemoteAddr: %s", r.Method, r.URL.Path, r.RemoteAddr)
}
