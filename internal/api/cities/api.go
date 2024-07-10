package cities

import (
	"context"
	"time"

	"github.com/yngk19/weatherapp/internal/model/domain"
)

type citiesService interface {
	GetCities(ctx context.Context) ([]domain.Town, error)
	GetByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error)
	GetByDate(ctx context.Context, date time.Time) (*domain.WeatherForecast, error)
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
