package api

import (
	"encoding/json"
	"net/http"

	"github.com/hriday111/weather-backend/internal/service"
)

func handleForecast(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon", http.StatusBadRequest)
		return
	}

	data, err := service.GetForecast(lat, lon)
	if err != nil {
		http.Error(w, "Failed to fetch forecast", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func handleSummary(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en"
	}
	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon", http.StatusBadRequest)
		return
	}

	summary, err := service.GetSummary(lat, lon, lang)
	if err != nil {
		http.Error(w, "Failed to get summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
