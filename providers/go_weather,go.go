package providers

import (
	"github.com/rendyspratama/weather-check/models"

	"encoding/json"
	"fmt"
	"net/http"
)

type GoWeatherAPI struct{}

func (g *GoWeatherAPI) GetWeather(city string) (*models.Weather, error) {
	url := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather models.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
