// Package main provides a WeatherService for retrieving weather data from an external API.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherFetcher defines an interface for retrieving weather data.
type WeatherProvider interface {
	FetchCityWeather(city string) (WeatherData, error)
}

// WeatherService defines a service for retrieving weather data from a specified API.
type WeatherService struct {
	apiKey  string // API key for authentication with the weather API
	baseURL string // Base URL for the weather API endpoint
}

// NewWeatherService initializes and returns a new WeatherService instance with the specified API key.
func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		apiKey:  apiKey,
		baseURL: "https://api.openweathermap.org/data/2.5/weather",
	}
}

// FetchCityWeather retrieves weather data for a given city by making a request to the weather API.
func (ws *WeatherService) FetchCityWeather(city string) (WeatherData, error) {
	url := ws.constructURL(city)
	res, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer res.Body.Close()

	// Validate the HTTP response
	if err := ws.validateResponse(res); err != nil {
		return WeatherData{}, err
	}

	// Parse the response JSON into WeatherData struct
	return ws.parseWeatherData(res)
}

// constructURL builds the request URL with the city name and API key.
func (ws *WeatherService) constructURL(city string) string {
	return fmt.Sprintf("%s?q=%s&appid=%s", ws.baseURL, city, ws.apiKey)
}

// validateResponse checks the HTTP response status code and returns an error if not 200 OK.
func (ws *WeatherService) validateResponse(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching weather data: status code %d", res.StatusCode)
	}
	return nil
}

// parseWeatherData decodes the JSON response body into a WeatherData struct.
func (ws *WeatherService) parseWeatherData(res *http.Response) (WeatherData, error) {
	var weatherData WeatherData
	if err := json.NewDecoder(res.Body).Decode(&weatherData); err != nil {
		return WeatherData{}, fmt.Errorf("error parsing weather data: %v", err)
	}
	return weatherData, nil
}
