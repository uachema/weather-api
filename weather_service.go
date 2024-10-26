// Package main provides a WeatherService for retrieving weather data from an external API.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherService defines a service for retrieving weather data from a specified API.
// It holds an API key for authentication and a base URL for the API endpoint.
type WeatherService struct {
	apiKey  string // API key for authentication with the weather API
	baseURL string // Base URL for the weather API endpoint
}

// NewWeatherService initializes and returns a new WeatherService instance with the specified API key.
// It sets a default base URL pointing to OpenWeather's API endpoint. This default can be overridden
// by providing a custom URL in the WeatherService struct directly if needed.
func NewWeatherService(apiKey string) *WeatherService {
	url := "https://api.openweathermap.org/data/2.5/weather"
	return &WeatherService{
		apiKey:  apiKey,
		baseURL: url,
	}
}

// FetchCityWeather retrieves weather data for a given city by sending a request to the weather API.
// It constructs the API URL with the city name and API key, then makes a GET request to the endpoint.
//
// Parameters:
// - city: the name of the city for which weather data is requested.
//
// Returns:
// - WeatherData: a struct containing parsed weather data if the request is successful.
// - error: an error if the request fails or if the response cannot be parsed.
//
// FetchCityWeather handles HTTP errors by returning an error if the response status code is not 200 OK.
// It also parses the JSON response into a WeatherData struct, returning it on success.
func (ws *WeatherService) FetchCityWeather(city string) (WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s", ws.baseURL, city, ws.apiKey)

	// Make an HTTP GET request to fetch the weather data for the specified city
	res, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err // Return error if the HTTP request fails
	}
	defer res.Body.Close() // Ensure the response body is closed after reading

	// Check if the response status code is OK (200)
	if res.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("error fetching weather data: status code %d", res.StatusCode)
	}

	// Declare a variable to hold the unmarshalled data
	var weatherData WeatherData
	// Decode the JSON response into the weatherData struct
	err = json.NewDecoder(res.Body).Decode(&weatherData)
	if err != nil {
		return WeatherData{}, err // Return error if unmarshalling fails
	}

	return weatherData, nil // Return the weather data if successful
}
