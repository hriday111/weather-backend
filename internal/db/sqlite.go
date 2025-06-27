package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hriday111/weather-backend/internal/config"
	"github.com/hriday111/weather-backend/internal/model"
)

var db *sql.DB

func InitDB() {
	var err error
	if _, err := os.Stat(config.SQLiteDBFile); os.IsNotExist(err) {
		file, _ := os.Create(config.SQLiteDBFile)
		file.Close()
	}

	db, err = sql.Open("sqlite3", config.SQLiteDBFile)
	if err != nil {
		log.Fatalf("failed to open sqlite db: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS forecast_cache (
		key TEXT PRIMARY KEY,
		data TEXT,
		timestamp DATETIME
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("failed to create forecast_cache table: %v", err)
	}
}

func GetForecastFromCache(key string, ttl time.Duration) (*model.ForecastResponse, bool) {
	var dataStr string
	var ts time.Time

	row := db.QueryRow("SELECT data, timestamp FROM forecast_cache WHERE key = ?", key)
	if err := row.Scan(&dataStr, &ts); err != nil {
		return nil, false
	}

	if time.Since(ts) > ttl {
		_ = DeleteForecast(key)
		return nil, false
	}

	var forecast model.ForecastResponse
	err := json.Unmarshal([]byte(dataStr), &forecast)
	if err != nil {
		return nil, false
	}
	return &forecast, true
}

func SetForecastToCache(key string, forecast *model.ForecastResponse) error {
	jsonBytes, err := json.Marshal(forecast)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT OR REPLACE INTO forecast_cache (key, data, timestamp)
		VALUES (?, ?, ?);
	`, key, string(jsonBytes), time.Now())

	return err
}

func DeleteForecast(key string) error {
	_, err := db.Exec("DELETE FROM forecast_cache WHERE key = ?", key)
	return err
}
