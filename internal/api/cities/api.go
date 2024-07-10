package cities

import (
	"context"
	"time"

	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
)

type citiesService interface {
	GetCities(ctx context.Context) ([]domain.Town, error)
	GetForecastByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error)
	GetForecastByDate(ctx context.Context, date time.Time) (*domain.WeatherForecast, error)
	GetShortByCityID(ctx context.Context, id int) (*dto.ShortForecast, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type API struct {
	logger  logger
	service citiesService
}

func NewAPI(logger logger, service citiesService) *API {
	return &API{
		logger:  logger,
		service: service,
	}
}
