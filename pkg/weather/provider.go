package weather

// WeatherProvider defines an interface for retrieving weather data.
type WeatherProvider interface {
	FetchCityWeather(city string) (WeatherData, error)         // Fetch weather data for a single city
	FetchCitiesWeather(cities []string) ([]WeatherData, error) // Fetch weather data for multiple cities
}
