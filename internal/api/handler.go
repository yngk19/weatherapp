package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yngk19/weatherapp/internal/api/cities"
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

func New(logger logger, citcitiesService citiesService) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	citiesAPI := cities.NewAPI(logger, citcitiesService)
	api := r.Group("/")
	api.GET("/cities", citiesAPI.GetCities)

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
