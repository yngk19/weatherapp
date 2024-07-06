package forecastsrepo

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yngk19/weatherapp/internal/model/domain"
	"github.com/yngk19/weatherapp/internal/model/dto"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) Create(ctx context.Context, forecast dto.WeatherForecast, city domain.Town) error {
	var cityName string
	query := `
		INSERT INTO forecasts (temp, city_name, prediction_date, data)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (city_name) DO UPDATE
		SET prediction_date = EXCLUDED.prediction_date,
		temp = EXCLUDED.temp,
		data = EXCLUDED.data
		RETURNING city_name;
	`
	if err := r.pool.QueryRow(ctx, query, forecast.List[0].Main.Temp, city.Name, forecast.List[0].DtTxt, forecast.List).Scan(&cityName); err != nil {
		return fmt.Errorf("repository.Cities.Create: %w", err)
	}
	fmt.Println(cityName)
	return nil
}

/*
func (r *Repo) GetAll(ctx context.Context) ([]domain.Town, error) {
	var towns []domain.Town
	query := `
		SELECT c.id, c.name, c.country, c.lat, c.lon
		FROM cities c;
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Cities.GetAll: %w", err)
	}
	for rows.Next() {
		town := domain.Town{}
		err := rows.Scan(
			&town.ID,
			&town.Name,
			&town.Country,
			&town.Lat,
			&town.Lon,
		)
		if err != nil {
			return nil, fmt.Errorf("repository.Cities.GetAll: %w", err)
		}
		towns = append(towns, town)
	}
	return towns, nil
}
*/
