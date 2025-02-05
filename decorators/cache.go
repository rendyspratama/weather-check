package decorators

import (
	"github.com/rendyspratama/weather-check/models"
	"github.com/rendyspratama/weather-check/providers"

	"sync"
	"time"
)

type CachedWeatherProvider struct {
	Provider providers.WeatherProvider
	Cache    map[string]*models.Weather
	Mutex    sync.Mutex
	TTL      time.Duration
}

func (c *CachedWeatherProvider) GetWeather(city string) (*models.Weather, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if data, exists := c.Cache[city]; exists {
		return data, nil
	}

	weather, err := c.Provider.GetWeather(city)
	if err != nil {
		return nil, err
	}

	c.Cache[city] = weather
	go func() {
		time.Sleep(c.TTL)
		delete(c.Cache, city)
	}()

	return weather, nil
}
