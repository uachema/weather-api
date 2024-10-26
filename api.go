// Package main implements a simple HTTP API server with structured response handling and error handling.
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer defines an HTTP API server with a specified listening address, router for routing requests,
// and an injected WeatherService for fetching weather data.
type APIServer struct {
	listenAddr     string         // Server's listening address
	router         *mux.Router    // Router for managing routes and handling requests
	weatherService WeatherService // Service for fetching weather data
}

// NewAPIServer creates and initializes a new APIServer with the given listening address and WeatherService.
// It sets up the router with predefined routes and returns the server instance.
func NewAPIServer(listenAddr string, weatherService WeatherService) *APIServer {
	server := &APIServer{
		listenAddr:     listenAddr,
		router:         mux.NewRouter(),
		weatherService: weatherService,
	}
	server.setupRoutes()
	return server
}

// setupRoutes defines the routes for the APIServer and associates each route with a handler function.
// The root endpoint is mapped to handleRoot.
func (s *APIServer) setupRoutes() {
	s.router.HandleFunc("/", makeHTTPHandlerFunc(s.handleRoot)).Methods("GET")
}

// Run starts the APIServer on the specified listening address and begins serving requests.
// It logs the server status and returns an error if the server fails to start.
func (s *APIServer) Run() error {
	log.Printf("Server started on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.router)
}

// handleRoot handles requests to the root endpoint ("/").
// It expects a query parameter "city" in the request URL, allowing multiple values (e.g., ?city=lahore&city=karachi).
// For each city specified, it fetches the weather data using the WeatherService and returns it as a JSON response.
func (s *APIServer) handleRoot(w http.ResponseWriter, r *http.Request) error {
	logRequest(r)

	// Get the "city" query parameter from the request URL
	cities := r.URL.Query()["city"]
	if len(cities) == 0 {
		// If no cities are provided, return a bad request response
		return writeJSON(w, http.StatusBadRequest, APIError{Error: "At least one city parameter is required"})
	}

	results := make([]WeatherData, 0, len(cities))

	// Fetch weather data for each specified city
	for _, city := range cities {
		weatherData, err := s.weatherService.FetchCityWeather(city)
		if err == nil {
			results = append(results, weatherData)
		}
	}

	v := APIResponse{Message: "Weather fetched successfully", Data: results}
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
// It sets the "Content-Type" header to "application/json" and returns any encoding error encountered.
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// logRequest logs details about an incoming HTTP request, including method, path, and remote IP address.
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
