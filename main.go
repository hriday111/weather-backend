package main

import (
	"log"
	"net/http"

	"github.com/hriday111/weather-backend/internal/api"
	"github.com/hriday111/weather-backend/internal/db"
)

func main() {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux)
	db.InitDB()

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
