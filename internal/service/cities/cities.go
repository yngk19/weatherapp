package cities

import (
	"context"
	"strings"

	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
	"github.com/yngk19/weatherapp/internal/service"
)

type citiesRepo interface {
	Create(ctx context.Context, city dto.Town) error
	GetAll(ctx context.Context) ([]domain.Town, error)
	GetByCityID(ctx context.Context, id int) (*domain.Town, error)
}

type forecastsRepo interface {
	Create(ctx context.Context, forecast dto.WeatherForecast, city domain.Town) error
	GetByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error)
	GetByDate(ctx context.Context, date string, id int) (*domain.WeatherForecast, error)
}

type Service struct {
	citiesRepo    citiesRepo
	forecastsRepo forecastsRepo
}

func New(citiesRepo citiesRepo, forecastsRepo forecastsRepo) *Service {
	return &Service{
		citiesRepo:    citiesRepo,
		forecastsRepo: forecastsRepo,
	}
}

func (s *Service) GetCities(ctx context.Context) ([]domain.Town, error) {
	cities, err := s.citiesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if cities == nil {
		return nil, service.ErrCitiesListIsEmpty
	}
	return cities, nil
}

func (s *Service) GetForecastByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error) {
	forecasts, err := s.forecastsRepo.GetByCityID(ctx, id)
	if err != nil {
		return nil, err
	}
	if forecasts == nil {
		return nil, service.ErrNoForecasts
	}
	return forecasts, nil
}

func (s *Service) GetForecastByDate(ctx context.Context, date string, id int) (*domain.WeatherForecast, error) {
	forecast, err := s.forecastsRepo.GetByDate(ctx, date, id)
	if err != nil {
		return nil, err
	}
	if forecast == nil {
		return nil, service.ErrNoForecastForThisDate
	}
	return forecast, nil
}

func (s *Service) GetShortByCityID(ctx context.Context, id int) (*dto.ShortForecast, error) {
	forecasts, err := s.forecastsRepo.GetByCityID(ctx, id)
	if err != nil {
		return nil, err
	}
	if forecasts == nil {
		return nil, service.ErrNoForecasts
	}
	city, err := s.citiesRepo.GetByCityID(ctx, id)
	if err != nil {
		return nil, err
	}
	if forecasts == nil {
		return nil, service.ErrNoForecasts
	}
	var averageTemp float32
	var dates []string
	for _, forecast := range forecasts {
		averageTemp += float32(forecast.Temperature)
		date := strings.Split(forecast.Date.String(), " ")[0]
		dates = append(dates, date)
	}
	averageTemp = averageTemp / float32(len(forecasts))
	shortForecast := dto.ShortForecast{
		Country:            city.Country,
		Name:               city.Name,
		Lat:                city.Lat,
		Lon:                city.Lon,
		AverageTemperature: float64(averageTemp),
		Dates:              dates,
	}
	return &shortForecast, nil
}
