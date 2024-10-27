// Package main implements a simple HTTP API server with structured response handling and error management.
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer defines an HTTP API server with a specified listening address, router for routing requests,
// and an injected WeatherProvider for fetching weather data.
type APIServer struct {
	listenAddr      string          // Server's listening address
	router          *mux.Router     // Router for managing routes and handling requests
	weatherProvider WeatherProvider // Provider for fetching weather data
}

// NewAPIServer creates and initializes a new APIServer with the given listening address and WeatherProvider.
// It sets up the router with predefined routes and returns the server instance.
//
// Parameters:
// - listenAddr: The address on which the server will listen for incoming requests.
// - weatherProvider: An implementation of WeatherProvider for fetching weather data.
//
// Returns:
// - *APIServer: The initialized APIServer instance.
func NewAPIServer(listenAddr string, weatherProvider WeatherProvider) *APIServer {
	server := &APIServer{
		listenAddr:      listenAddr,
		router:          mux.NewRouter(),
		weatherProvider: weatherProvider,
	}
	server.setupRoutes()
	return server
}

// setupRoutes defines the routes for the APIServer and associates each route with a handler function.
func (s *APIServer) setupRoutes() {
	s.router.HandleFunc("/", makeHTTPHandlerFunc(s.handleRoot)).Methods("GET")
}

// Run starts the APIServer on the specified listening address and begins serving requests.
// It logs the server status and returns an error if the server fails to start.
//
// Returns:
// - error: An error if the server fails to start.
func (s *APIServer) Run() error {
	log.Printf("Server started on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.router)
}

// handleRoot handles requests to the root endpoint ("/").
// It expects a query parameter "city" in the request URL, allowing multiple values (e.g., ?city=lahore&city=karachi).
// For each city specified, it fetches the weather data using the WeatherProvider and returns it as a JSON response.
//
// Parameters:
// - w: The http.ResponseWriter for writing the response.
// - r: The http.Request containing the request data.
//
// Returns:
// - error: An error if the handler fails, managed by makeHTTPHandlerFunc.
func (s *APIServer) handleRoot(w http.ResponseWriter, r *http.Request) error {
	logRequest(r)

	// Get the "city" query parameter from the request URL
	cities := r.URL.Query()["city"]
	if len(cities) == 0 {
		// If no cities are provided, return a bad request response
		return writeJSON(w, http.StatusBadRequest, APIError{Error: "At least one city parameter is required"})
	}

	// Fetch weather data for the specified cities
	results, err := s.weatherProvider.FetchCitiesWeather(cities)
	if err != nil {
		// If an error occurs during fetching, log it and return an internal server error
		log.Printf("Error fetching weather data: %v", err)
		return writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	// Create a successful response
	response := APIResponse{Message: "Weather fetched successfully", Data: results}
	return writeJSON(w, http.StatusOK, response)
}

// handleFunc defines a function type that processes HTTP requests and returns an error if one occurs.
type handleFunc func(w http.ResponseWriter, r *http.Request) error

// makeHTTPHandlerFunc wraps a handleFunc, converting it to an http.HandlerFunc.
// It manages error handling by responding with a JSON error message if the handler function returns an error.
//
// Parameters:
// - f: The handleFunc to wrap.
//
// Returns:
// - http.HandlerFunc: The wrapped handler function.
func makeHTTPHandlerFunc(f handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("HTTP handler error: %v", err)
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

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
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// logRequest logs details about an incoming HTTP request, including method, path, and remote IP address.
//
// Parameters:
// - r: The http.Request to log.
func logRequest(r *http.Request) {
	log.Printf("Request received - Method: %s, Path: %s, RemoteAddr: %s", r.Method, r.URL.Path, r.RemoteAddr)
}

// APIResponse represents a successful JSON response containing a message and optional data.
type APIResponse struct {
	Message string `json:"message"` // The response message to be returned in JSON format
	Data    any    `json:"data"`    // Optional data to be returned in JSON format
}

// APIError represents an error response in JSON format.
type APIError struct {
	Error string `json:"error"` // The error message to be returned in JSON format
}
