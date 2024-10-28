// Package server implements a simple HTTP API server with structured response handling and error management.
package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uachema/weather-api/pkg/util"
	"github.com/uachema/weather-api/pkg/weather"
)

// APIServer defines an HTTP API server with a specified listening address, router for routing requests,
// and an injected WeatherProvider for fetching weather data.
type APIServer struct {
	listenAddr      string                  // Server's listening address
	router          *mux.Router             // Router for managing routes and handling requests
	weatherProvider weather.WeatherProvider // Provider for fetching weather data
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
func NewAPIServer(listenAddr string, weatherProvider weather.WeatherProvider) *APIServer {
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
	if err := http.ListenAndServe(s.listenAddr, s.router); err != nil {
		log.Fatalf("Could not start server: %v", err) // Log fatal error if server fails to start
		return err
	}
	return nil
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
			util.WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}
