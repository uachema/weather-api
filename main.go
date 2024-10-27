// Package main initializes and runs an API server that fetches weather data from an external service.
// It uses environment variables for configuration, such as the server's port and the OpenWeather API key.
package main

import (
	"fmt"
	"log"
	"os"
)

// main function sets up the necessary configurations for the server,
// including retrieving environment variables and initializing the weather service.
// It then starts the server and listens on the specified port.
func main() {
	// Retrieve the port on which the server will listen from the PORT environment variable.
	// If PORT is not set, use the default port "3000".
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default to port 3000 if not specified
	}

	// Retrieve the OpenWeather API key from the environment variables.
	// This key is required to authenticate requests to the weather service.
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY environment variable is required") // Log an error and exit if the API key is missing
	}

	// Initialize a new instance of the WeatherService using the API key.
	weatherService := NewWeatherService(apiKey)

	// Construct the server's listening address by combining ":" with the specified port.
	listenAddr := fmt.Sprintf(":%s", port)

	// Initialize a new API server with the specified listening address and weather service instance.
	server := NewAPIServer(listenAddr, weatherService)

	// Start the server and log any errors that occur if the server fails to start.
	if err := server.Run(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
