package main

// WeatherData represents the weather information returned from the OpenWeather API.
// It includes data for coordinates, current conditions, main metrics like temperature, wind, and visibility,
// as well as information about clouds, system data, and time-based details.
type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"` // Longitude of the city location
		Lat float64 `json:"lat"` // Latitude of the city location
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`          // Weather condition ID
		Main        string `json:"main"`        // Group of weather parameters (e.g., Rain, Snow, etc.)
		Description string `json:"description"` // Description of the weather condition
		Icon        string `json:"icon"`        // Icon ID for visual representation of the weather
	} `json:"weather"`

	Base string `json:"base"` // Internal parameter used by the API

	Main struct {
		Temp      float64 `json:"temp"`       // Current temperature
		FeelsLike float64 `json:"feels_like"` // Perceived temperature
		TempMin   float64 `json:"temp_min"`   // Minimum temperature at the moment
		TempMax   float64 `json:"temp_max"`   // Maximum temperature at the moment
		Pressure  int     `json:"pressure"`   // Atmospheric pressure
		Humidity  int     `json:"humidity"`   // Humidity percentage
		SeaLevel  int     `json:"sea_level"`  // Sea level atmospheric pressure (if available)
		GrndLevel int     `json:"grnd_level"` // Ground level atmospheric pressure (if available)
	} `json:"main"`

	Visibility int `json:"visibility"` // Visibility distance in meters

	Wind struct {
		Speed float64 `json:"speed"` // Wind speed
		Deg   int     `json:"deg"`   // Wind direction in degrees
		Gust  float64 `json:"gust"`  // Wind gust speed (if available)
	} `json:"wind"`

	Rain struct {
		OneH float64 `json:"1h"` // Rain volume in the last hour (if available)
	} `json:"rain"`

	Clouds struct {
		All int `json:"all"` // Cloudiness percentage
	} `json:"clouds"`

	Dt int `json:"dt"` // Time of data calculation (Unix timestamp)

	Sys struct {
		Type    int    `json:"type"`    // Internal parameter
		ID      int    `json:"id"`      // Internal parameter
		Country string `json:"country"` // Country code (e.g., "US")
		Sunrise int    `json:"sunrise"` // Sunrise time (Unix timestamp)
		Sunset  int    `json:"sunset"`  // Sunset time (Unix timestamp)
	} `json:"sys"`

	Timezone int    `json:"timezone"` // Shift in seconds from UTC
	ID       int    `json:"id"`       // City ID
	Name     string `json:"name"`     // City name
	Cod      int    `json:"cod"`      // HTTP status code of the response
}
