// Package main implements a simple HTTP API server with structured response handling and error handling.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer defines an HTTP API server with a specified listening address and router for routing requests.
type APIServer struct {
	listenAddr string      // Server's listening address
	router     *mux.Router // Router for managing routes and handling requests
}

// NewAPIServer creates and initializes a new APIServer with the given listening address.
// It sets up the router with predefined routes and returns the server instance.
func NewAPIServer(listenAddr string) *APIServer {
	server := &APIServer{
		listenAddr: listenAddr,
		router:     mux.NewRouter(),
	}
	server.setupRoutes()
	return server
}

// setupRoutes defines the routes for the APIServer and associates each route with a handler function.
func (s *APIServer) setupRoutes() {
	// Define route for the root endpoint
	s.router.HandleFunc("/", makeHTTPHandlerFunc(s.handleRoot)).Methods("GET")

	// Define route for the dynamic ID endpoint
	s.router.HandleFunc("/{id}", makeHTTPHandlerFunc(s.handleRootID)).Methods("GET")
}

// Run starts the APIServer on the specified listening address and begins serving requests.
// It logs the server status and returns an error if the server fails to start.
func (s *APIServer) Run() error {
	log.Printf("Server started on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.router)
}

// handleRoot handles requests to the root endpoint ("/").
// It responds with a JSON message containing "Hello World".
func (s *APIServer) handleRoot(w http.ResponseWriter, r *http.Request) error {
	logRequest(r)
	v := APIResponse{Message: "Hello World"}
	return writeJSON(w, http.StatusOK, v)
}

// handleRootID handles requests to the endpoint with a dynamic ID ("/{id}").
// It responds with a JSON message containing "Hello World" and the provided ID.
func (s *APIServer) handleRootID(w http.ResponseWriter, r *http.Request) error {
	logRequest(r)
	vars := mux.Vars(r)
	id := vars["id"]
	v := APIResponse{Message: fmt.Sprintf("Hello World %s", id)}
	return writeJSON(w, http.StatusOK, v)
}

// handleFunc defines a function type that processes HTTP requests and returns an error if one occurs.
type handleFunc func(w http.ResponseWriter, r *http.Request) error

// makeHTTPHandlerFunc wraps a handleFunc, converting it to an http.HandlerFunc.
// It manages error handling by responding with a JSON error message if the handler function returns an error.
func makeHTTPHandlerFunc(f handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("HTTP handler error: %v", err)
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

// writeJSON encodes a given value as JSON and writes it to the http.ResponseWriter with the specified HTTP status code.
// It sets the "Content-Type" header to "application/json".
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// logRequest logs details about an incoming HTTP request, including method, path, and remote IP address.
func logRequest(r *http.Request) {
	log.Printf("Request received - Method: %s, Path: %s, RemoteAddr: %s", r.Method, r.URL.Path, r.RemoteAddr)
}

// APIResponse represents a successful JSON response containing a message.
type APIResponse struct {
	Message string `json:"message"` // The response message to be returned in JSON format
}

// APIError represents an error response in JSON format.
type APIError struct {
	Error string `json:"error"` // The error message to be returned in JSON format
}
