package api

import (
	"encoding/json"
	"net/http"

	"github.com/hriday111/weather-backend/internal/service"
)

// This file contains HTTP handlers for weather-related data.

// handleForecast handles requests for weather forecast data.
func handleForecast(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")   // Latitude parameter
	lon := r.URL.Query().Get("lon")   // Longitude parameter
	lang := r.URL.Query().Get("lang") // Language parameter
	if lang == "" {
		lang = "en"
	}
	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon", http.StatusBadRequest)
		return
	}

	data, err := service.GetForecast(lat, lon, lang) // Fetch forecast data
	if err != nil {
		http.Error(w, "Failed to fetch forecast", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data) // Send response as JSON
}

// handleSummary handles requests for weather summary data.
func handleSummary(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")   // Latitude parameter
	lon := r.URL.Query().Get("lon")   // Longitude parameter
	lang := r.URL.Query().Get("lang") // Language parameter
	if lang == "" {
		lang = "en"
	}
	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon", http.StatusBadRequest)
		return
	}

	summary, err := service.GetSummary(lat, lon, lang) // Fetch summary data
	if err != nil {
		http.Error(w, "Failed to get summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary) // Send response as JSON
}
