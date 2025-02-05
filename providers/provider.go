package providers

import "github.com/rendyspratama/weather-check/models"

type WeatherProvider interface {
	GetWeather(city string) (*models.Weather, error)
}

func NewWeatherProvider(providerType string) WeatherProvider {
	switch providerType {
	case "goWeather":
		return &GoWeatherAPI{}
	default:
		return &GoWeatherAPI{} // Default provider
	}
}
