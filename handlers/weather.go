package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// WeatherResponse defines the response structure from the API
type WeatherResponse struct {
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
	Description string `json:"description"`
}

// GetWeather fetches weather data for a city
func GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := strings.TrimSpace(vars["city"])

	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", city)
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "City not found", http.StatusNotFound)
		return
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
