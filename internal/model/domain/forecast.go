package domain

import (
	"time"

	"github.com/yngk19/weatherapp/internal/model/dto"
)

type WeatherForecast struct {
	ID          int
	Temperature float64
	Date        time.Time
	FullInfo
}

type FullInfo struct {
	dto.Main   `json:"main"`
	Weather    []dto.Weather
	dto.Clouds `json:"clouds"`
	dto.Wind   `json:"wind"`
	Visibility int     `json:"visibility"`
	Pop        float64 `json:"pop"`
	dto.Sys    `json:"sys"`
}
