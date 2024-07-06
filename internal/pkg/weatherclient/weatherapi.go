package weatherclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
)

type ForecastInterface interface {
	Create(context.Context, dto.WeatherForecast, domain.Town) error
}

type CityInterface interface {
	GetAll(context.Context) ([]domain.Town, error)
}

func GetForecasts(forecastRepo ForecastInterface, cityRepo CityInterface, apiToken string) error {
	client := http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: 100}}
	var wg sync.WaitGroup
	cities, err := cityRepo.GetAll(context.Background())
	if err != nil {
		return fmt.Errorf("weatherapigo.GetForecasts: %w", err)
	}
	for _, city := range cities {
		wg.Add(1)
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s", city.Lat, city.Lon, apiToken)
		go func(url string) {
			var forecast dto.WeatherForecast
			defer wg.Done()
			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("%s: %s\n", url, err)
				return
			}
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
				fmt.Printf("%s: %s\n", url, err)
				return
			}
			err = forecastRepo.Create(context.Background(), forecast, city)
			if err != nil {
				fmt.Printf("%s: %s\n", url, err)
				return
			}
		}(url)
	}

	wg.Wait()
	return nil
}
