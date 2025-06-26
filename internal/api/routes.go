package api

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/forecast", handleForecast)
	mux.HandleFunc("/summary", handleSummary)
}
