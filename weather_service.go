// Package main provides a WeatherService for retrieving weather data from an external API.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// WeatherProvider defines an interface for retrieving weather data.
type WeatherProvider interface {
	FetchCityWeather(city string) (WeatherData, error)         // Fetch weather data for a single city
	FetchCitiesWeather(cities []string) ([]WeatherData, error) // Fetch weather data for multiple cities
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
// It constructs the API URL, sends a GET request, and parses the response JSON into a WeatherData struct.
//
// Parameters:
// - city: The name of the city for which to fetch weather data.
//
// Returns:
// - WeatherData: The parsed weather data for the city.
// - error: An error if the request fails or if the response cannot be parsed.
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

// FetchCitiesWeather retrieves weather data for multiple cities concurrently.
// It launches a goroutine for each city to fetch its weather data and collects the results.
//
// Parameters:
// - cities: A slice of city names for which to fetch weather data.
//
// Returns:
// - []WeatherData: A slice containing the parsed weather data for each city.
// - error: An error if any request fails, but results may still be returned for successfully fetched cities.
func (ws *WeatherService) FetchCitiesWeather(cities []string) ([]WeatherData, error) {
	var wg sync.WaitGroup
	ch := make(chan WeatherData, len(cities)) // Buffered channel to hold results

	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()

			// Fetch weather data for the city
			url := ws.constructURL(city)
			res, err := http.Get(url)
			if err != nil {
				log.Printf("Error fetching weather data for city %s: %v", city, err)
				return // Log error and return
			}
			defer res.Body.Close()

			// Validate and parse the response
			if err := ws.validateResponse(res); err != nil {
				log.Printf("Invalid response for city %s: %v", city, err)
				return // Log error and return
			}

			weatherData, err := ws.parseWeatherData(res)
			if err != nil {
				log.Printf("Error parsing weather data for city %s: %v", city, err)
				return // Log error and return
			}
			ch <- weatherData // Send the weather data to the channel
		}(city)
	}

	// Close the channel after all goroutines have completed
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the results channel
	}()

	// Collect results from the channel
	var result []WeatherData
	for data := range ch {
		result = append(result, data)
	}

	return result, nil
}

// constructURL builds the request URL with the city name and API key.
// This method constructs a complete API URL for the specified city.
//
// Parameters:
// - city: The name of the city to construct the URL for.
//
// Returns:
// - string: The constructed URL for fetching weather data for the city.
func (ws *WeatherService) constructURL(city string) string {
	return fmt.Sprintf("%s?q=%s&appid=%s", ws.baseURL, city, ws.apiKey)
}

// validateResponse checks the HTTP response status code and returns an error if it is not 200 OK.
// It ensures that the API request was successful before proceeding to parse the response.
//
// Parameters:
// - res: The HTTP response to validate.
//
// Returns:
// - error: An error if the response status code indicates failure (not 200).
func (ws *WeatherService) validateResponse(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching weather data: status code %d", res.StatusCode)
	}
	return nil
}

// parseWeatherData decodes the JSON response body into a WeatherData struct.
// This method reads the response body and unmarshals the JSON data into a predefined struct.
//
// Parameters:
// - res: The HTTP response containing the weather data.
//
// Returns:
// - WeatherData: The parsed weather data.
// - error: An error if unmarshalling the JSON fails.
func (ws *WeatherService) parseWeatherData(res *http.Response) (WeatherData, error) {
	var weatherData WeatherData
	if err := json.NewDecoder(res.Body).Decode(&weatherData); err != nil {
		return WeatherData{}, fmt.Errorf("error parsing weather data: %v", err)
	}
	return weatherData, nil
}
