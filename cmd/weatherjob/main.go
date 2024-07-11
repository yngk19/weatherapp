package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"github.com/yngk19/weatherapp/internal/config"
	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
	citiesrepo "github.com/yngk19/weatherapp/internal/repository/cities"
	forecastsrepo "github.com/yngk19/weatherapp/internal/repository/forecasts"
	"github.com/yngk19/weatherapp/pkg/db"
	"github.com/yngk19/weatherapp/pkg/weatherclient"
)

type citiesRepo interface {
	GetAll(ctx context.Context) ([]domain.Town, error)
}

type forecastsRepo interface {
	Create(ctx context.Context, forecast dto.WeatherForecast, city domain.Town) error
}

type WeatherGetJob struct {
	APIKey        string
	citiesRepo    citiesRepo
	forecastsRepo forecastsRepo
}

func main() {
	cfg := config.MustLoad()
	location, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	citiesRepo := citiesrepo.New(pool)
	forecastsRepo := forecastsrepo.New(pool)
	weatherJob := WeatherGetJob{
		APIKey:        cfg.OpenWeatherAPIKey,
		citiesRepo:    citiesRepo,
		forecastsRepo: forecastsRepo,
	}
	cronJob := cron.NewWithLocation(location)
	cronJob.AddJob("*/5 * * * *", weatherJob)
	go func() {
		fmt.Println("Starting GetForecasts() job")
		cronJob.Run()
	}()
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	<-signals
	fmt.Println("Cron scheduler is stopped!")
	cronJob.Stop()
}

func (job WeatherGetJob) Run() {
	err := weatherclient.GetForecasts(job.forecastsRepo, job.citiesRepo, job.APIKey)
	if err != nil {
		log.Fatalln(err)
	}
}
