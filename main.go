package main

import (
	"fmt"
	"log"
	"os"
)

// main function initializes the server, setting up environment variables for port and API key.
func main() {
	// Get the PORT environment variable or use the default value of "3000"
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Get the OpenWeather API key from environment variables
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY environment variable is required")
	}

	// Initialize the WeatherService with the API key
	weatherService := NewWeatherService(apiKey)

	// Define the server's listening address using the port
	listenAddr := fmt.Sprintf(":%s", port)

	// Create and start the API server with the specified address and weather service
	server := NewAPIServer(listenAddr, *weatherService)
	if err := server.Run(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
