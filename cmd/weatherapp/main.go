package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/yngk19/weatherapp/internal/api"
	"github.com/yngk19/weatherapp/internal/config"
	"github.com/yngk19/weatherapp/internal/pkg/db"
	"github.com/yngk19/weatherapp/internal/pkg/geoapiclient"
	"github.com/yngk19/weatherapp/internal/pkg/logger"
	"github.com/yngk19/weatherapp/internal/pkg/weatherclient"
	citiesrepo "github.com/yngk19/weatherapp/internal/repository/cities"
	forecastsrepo "github.com/yngk19/weatherapp/internal/repository/forecasts"
	"github.com/yngk19/weatherapp/internal/service/cities"
)

func main() {
	cfg := config.MustLoad()
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	myLogger, err := logger.New(cfg.LogFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	citiesRepo := citiesrepo.New(pool)
	forecastsRepo := forecastsrepo.New(pool)
	citiesService := cities.New(citiesRepo, forecastsRepo)
	err = geoapiclient.GetCities(citiesRepo, cfg.OpenWeatherAPIKey)
	if err != nil {
		log.Fatalln(err)
	}
	err = weatherclient.GetForecasts(forecastsRepo, citiesRepo, cfg.OpenWeatherAPIKey)
	if err != nil {
		log.Fatalln(err)
	}
	r, err := api.New(
		myLogger,
		citiesService,
	)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Failed to start http server!")
		}
		fmt.Println("Listening on localhost:3000")
	}()
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	<-signals
	fmt.Println("Server stoped!")
	srv.Shutdown(context.Background())
}
