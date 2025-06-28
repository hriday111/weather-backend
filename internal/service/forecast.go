package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/hriday111/weather-backend/internal/config"
	"github.com/hriday111/weather-backend/internal/db"
	"github.com/hriday111/weather-backend/internal/model"
	"github.com/hriday111/weather-backend/internal/util"
)

func GetForecast(lat, lon, lang string) (*model.ForecastResponse, error) {
	key := fmt.Sprintf("%s,%s", lat, lon)
	cached, ok := db.GetForecastFromCache(key, 30*time.Minute)
	if ok {
		go func() {
			fresh, err := fetchAndBuildForecast(lat, lon, lang)
			if err == nil {
				_ = db.SetForecastToCache(key, fresh)
			}
		}()

		return cached, nil
	}

	fresh, err := fetchAndBuildForecast(lat, lon, lang)
	if err != nil {
		return nil, err
	}
	_ = db.SetForecastToCache(key, fresh)
	return fresh, nil

}

func fetchAndBuildForecast(lat, lon, lang string) (*model.ForecastResponse, error) {
	url := fmt.Sprintf("%s?latitude=%s&longitude=%s&daily=temperature_2m_max,temperature_2m_min,weathercode,sunshine_duration&hourly=pressure_msl&timezone=auto", config.OpenMeteoBaseURL, lat, lon) //fmt.Println("Requesting URL:", url)

	resp, err_get := http.Get(url)
	if err_get != nil {
		return nil, fmt.Errorf("failed to call weather API: %w", err_get)
	}
	defer resp.Body.Close()

	var apiData model.APIResponse
	err_dec := json.NewDecoder(resp.Body).Decode(&apiData)
	if err_dec != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err_dec)
	}

	var forecast model.ForecastResponse
	forecast.PressureReadings = apiData.Hourly.PressureMSL

	for i := range apiData.Daily.Time {
		// inside the loop
		parsedDate, err := time.Parse("2006-01-02", apiData.Daily.Time[i])
		if err != nil {
			parsedDate = time.Now() // fallback
		}

		dayOfWeek := parsedDate.Weekday().String()

		energy := config.InstallationPowerKW * (apiData.Daily.SunshineDuration[i] / 60) * config.PanelEfficiency
		forecast.Days = append(forecast.Days, model.ForecastDay{
			Date:        apiData.Daily.Time[i],
			Day:         util.Translate(dayOfWeek, lang),
			WeatherCode: apiData.Daily.WeatherCode[i],
			TempMax:     apiData.Daily.TemperatureMax[i],
			TempMin:     apiData.Daily.TemperatureMin[i],
			EnergyKWh:   energy,
		})
	}
	return &forecast, nil
}
