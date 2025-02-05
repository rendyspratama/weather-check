package services

import (
	"github.com/rendyspratama/weather-check/models"
	"github.com/rendyspratama/weather-check/providers"
)

type WeatherService struct {
	Provider providers.WeatherProvider
}

func (s *WeatherService) GetWeather(city string) (*models.Weather, error) {
	return s.Provider.GetWeather(city)
}
