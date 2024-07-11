package geoapiclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/yngk19/weatherapp/internal/model/dto"
)

var cities = []string{
	"Moscow",
	"Ufa",
	"Kazan",
	"Yekaterinburg",
	"Novosibirsk",
	"Tomsk",
	"Samara",
	"Petrozavodsk",
	"Perm",
	"Vladivostok",
	"Ussuriysk",
	"Sochi",
	"Obninsk",
	"Arzamas",
	"Abakan",
	"Chelyabinsk",
	"Kaliningrad",
	"Tyumen",
	"London",
	"Krasnoyarsk",
}

type CitiesRepo interface {
	Create(context.Context, dto.Town) error
}

func GetCities(repo CitiesRepo, apiToken string) error {
	client := http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: 100}}

	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, apiToken)
		go func(url string) {
			towns := make([]dto.Town, 1)
			defer wg.Done()
			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("%s: %s\n", url, err)
				return
			}
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&towns); err != nil {
				fmt.Printf("%s: %s\n", url, err)
				return
			}
			err = repo.Create(context.Background(), towns[0])
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return
				}
				fmt.Printf("geoapiclient.GetCities: %s - %s", towns[0].Name, err.Error())
				return
			}
		}(url)
	}

	wg.Wait()
	return nil
}
