package main

import (
	"context"
	"log"

	"github.com/yngk19/weatherapp/internal/config"
	"github.com/yngk19/weatherapp/internal/pkg/db"
	"github.com/yngk19/weatherapp/internal/pkg/geoapiclient"
	"github.com/yngk19/weatherapp/internal/pkg/weatherclient"
	citiesrepo "github.com/yngk19/weatherapp/internal/repository/cities"
	forecastsrepo "github.com/yngk19/weatherapp/internal/repository/forecasts"
)

func main() {
	cfg := config.MustLoad()
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	citiesRepo := citiesrepo.New(pool)
	forecastsRepo := forecastsrepo.New(pool)
	err = geoapiclient.GetCities(citiesRepo, cfg.OpenWeatherAPIKey)
	if err != nil {
		log.Fatalln(err)
	}
	err = weatherclient.GetForecasts(forecastsRepo, citiesRepo, cfg.OpenWeatherAPIKey)
	if err != nil {
		log.Fatalln(err)
	}
}
