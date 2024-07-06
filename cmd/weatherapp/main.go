package main

import (
	"context"
	"log"

	"github.com/yngk19/weatherapp/internal/config"
	"github.com/yngk19/weatherapp/internal/pkg/db"
	"github.com/yngk19/weatherapp/internal/pkg/geoapiclient"
	citiesrepo "github.com/yngk19/weatherapp/internal/repository/cities"
)

func main() {
	cfg := config.MustLoad()
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	citiesrepo := citiesrepo.New(pool)
	err = geoapiclient.GetCities(citiesrepo, cfg.OpenWeatherAPIKey)
	if err != nil {
		log.Fatalln(err)
	}
}
