package server

import (
	"log"
	"net/http"

	"github.com/uachema/weather-api/pkg/util"
)

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
	util.LogRequest(r)

	// Get the "city" query parameter from the request URL
	cities := r.URL.Query()["city"]
	if len(cities) == 0 {
		// If no cities are provided, return a bad request response
		log.Println("No cities provided in the query parameter") // Log the missing cities
		return util.WriteJSON(w, http.StatusBadRequest, APIError{Error: "At least one city parameter is required"})
	}

	// Fetch weather data for the specified cities
	results, err := s.weatherProvider.FetchCitiesWeather(cities)
	if err != nil {
		// If an error occurs during fetching, log it and return an internal server error
		log.Printf("Error fetching weather data: %v", err)
		return util.WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	// Create a successful response
	response := APIResponse{Message: "Weather fetched successfully", Data: results}
	log.Printf("Successfully fetched weather data for cities: %v", cities) // Log successful data fetch
	return util.WriteJSON(w, http.StatusOK, response)
}
