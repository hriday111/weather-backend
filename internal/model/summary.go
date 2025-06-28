package model

// Represents weekly weather summary data.
type WeeklySummary struct {
	AveragePressure float64 `json:"average_pressure"`
	AverageSunHours float64 `json:"average_sun_hours"`
	MinTemperature  float64 `json:"min_temperature"`
	MaxTemperature  float64 `json:"max_temperature"`
	WeekDescription string  `json:"week_description"` // "with precipitation" or "without precipitation"
}
