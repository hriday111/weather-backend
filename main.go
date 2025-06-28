package main

import (
	"log"
	"net/http"

	"github.com/hriday111/weather-backend/internal/api"
	"github.com/hriday111/weather-backend/internal/db"
)

// Entry point for the application.
func main() {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux)
	// WRAP your mux with EnableCORS
	wrapped := api.EnableCORS(mux)
	db.InitDB()

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", wrapped)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
