package main

import (
	"github.com/rendyspratama/weather-check/config"
	"github.com/rendyspratama/weather-check/routes"

	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Setup routes
	r := routes.SetupRoutes()

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
