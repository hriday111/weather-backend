package service

import (
	"math"

	"github.com/hriday111/weather-backend/internal/config"
	"github.com/hriday111/weather-backend/internal/model"
	"github.com/hriday111/weather-backend/internal/util"
	//"time"
)

func GetSummary(lat, lon, lang string) (*model.WeeklySummary, error) {
	forecast, err := GetForecast(lat, lon)
	if err != nil {
		return nil, err
	}

	var (
		sunSum      float64
		tempMin     = math.MaxFloat64
		tempMax     = -math.MaxFloat64
		rainyDays   int
		pressureSum float64
	)

	// Heuristic: weather codes 51-67, 80-99 usually indicate precipitation
	isRain := func(code int) bool {
		return (code >= 51 && code <= 67) || (code >= 80 && code <= 99)
	}

	for _, day := range forecast.Days {
		sunSum += day.EnergyKWh / (config.InstallationPowerKW * config.PanelEfficiency) // reverse-engineer sun hours from energy
		if day.TempMin < tempMin {
			tempMin = day.TempMin
		}
		if day.TempMax > tempMax {
			tempMax = day.TempMax
		}
		if isRain(day.WeatherCode) {
			rainyDays++
		}
	}
	for _, p := range forecast.PressureReadings {
		pressureSum += p
	}
	summary := &model.WeeklySummary{
		AveragePressure: pressureSum / float64(len(forecast.PressureReadings)),
		AverageSunHours: sunSum / float64(len(forecast.Days)),
		MinTemperature:  tempMin,
		MaxTemperature:  tempMax,
	}

	if rainyDays >= 4 {
		summary.WeekDescription = util.Translate("with_precipitation", lang)
	} else {
		summary.WeekDescription = util.Translate("without_precipitation", lang)
	}

	return summary, nil
}
