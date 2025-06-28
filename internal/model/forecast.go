package model

type APIResponse struct {
	Daily struct {
		Time             []string  `json:"time"`
		TemperatureMax   []float64 `json:"temperature_2m_max"`
		TemperatureMin   []float64 `json:"temperature_2m_min"`
		WeatherCode      []int     `json:"weathercode"`
		SunshineDuration []float64 `json:"sunshine_duration"`
	} `json:"daily"`
	Hourly struct {
		Time        []string  `json:"time"`
		PressureMSL []float64 `json:"pressure_msl"`
	} `json:"hourly"`
}
type ForecastDay struct {
	Date        string  `json:"date"`
	Day         string  `json:"day"`
	WeatherCode int     `json:"weather_code"`
	TempMax     float64 `json:"temp_max"`
	TempMin     float64 `json:"temp_min"`
	EnergyKWh   float64 `json:"energy_kwh"`
}

type ForecastResponse struct {
	Days             []ForecastDay `json:"days"`
	PressureReadings []float64     `json:"pressure_readings,omitempty"`
}
