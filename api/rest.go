package api

import (
	"github.com/rendyspratama/weather-check/decorators"
	"github.com/rendyspratama/weather-check/models"
	"github.com/rendyspratama/weather-check/providers"
	"github.com/rendyspratama/weather-check/services"

	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var weatherService = services.WeatherService{
	Provider: &decorators.CachedWeatherProvider{
		Provider: providers.NewWeatherProvider("goWeather"),
		Cache:    make(map[string]*models.Weather),
		TTL:      10 * time.Minute,
	},
}

func HandleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	city := strings.TrimPrefix(r.URL.Path, "/weather/")
	weather, err := weatherService.GetWeather(city)
	if err != nil {
		http.Error(w, "Failed to fetch weather", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
