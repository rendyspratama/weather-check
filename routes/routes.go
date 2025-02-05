package routes

import (
	"github.com/gorilla/mux"

	"github.com/rendyspratama/weather-check/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", handlers.GetWeather).Methods("GET")
	return r
}
