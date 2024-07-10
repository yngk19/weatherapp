package domain

import (
	"time"

	"github.com/yngk19/weatherapp/internal/model/dto"
)

type WeatherForecast struct {
	ID          int
	Temperature float64
	Date        time.Time
	DetailInfo  []dto.List
}
