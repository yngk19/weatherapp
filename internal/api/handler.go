package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yngk19/weatherapp/internal/api/cities"
	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
)

type citiesService interface {
	GetCities(ctx context.Context) ([]domain.Town, error)
	GetForecastByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error)
	GetForecastByDate(ctx context.Context, date string, id int) (*domain.WeatherForecast, error)
	GetShortByCityID(ctx context.Context, id int) (*dto.ShortForecast, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

func New(logger logger, citcitiesService citiesService) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	citiesAPI := cities.NewAPI(logger, citcitiesService)
	api := r.Group("/")
	api.GET("/cities", citiesAPI.GetCities)
	api.GET("/cities/:cityID/short", citiesAPI.GetShortForecast)
	api.GET("cities/:cityID/:date", citiesAPI.GetForecastByDate)

	return r, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}
