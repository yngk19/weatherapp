package forecastsrepo

import (
	"context"
	"fmt"
	"strings"

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
	query := `
		INSERT INTO forecasts (temp, city_id, predict_date, detail_info)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (predict_date, city_id) DO UPDATE
		SET temp = EXCLUDED.temp,
		detail_info = EXCLUDED.detail_info;
	`
	forecasts := make(map[string][]dto.List)
	for _, weather := range forecast.List {
		date := strings.Split(weather.DtTxt, " ")[0]
		forecasts[date] = append(forecasts[date], weather)
	}
	for date, weatherSet := range forecasts {
		fmt.Println(date)
		if err := r.pool.QueryRow(ctx, query, weatherSet[0].Main.Temp, city.ID, date, weatherSet).Scan(); err != nil {
			return fmt.Errorf("repository.Forecasts.Create: %w", err)
		}
	}
	return nil

}

func (r *Repo) GetByCityID(ctx context.Context) ([]domain.WeatherForecast, error) {
	var forecasts []domain.WeatherForecast
	query := `
		SELECT f.id, f.temp, f.predict_date, f.detail_info 
		FROM forecasts f 
		ORDER BY predict_date DESC
		LIMIT 5;
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Forecasts.GetAll: %w", err)
	}
	for rows.Next() {
		forecast := domain.WeatherForecast{}
		err := rows.Scan(
			&forecast.ID,
			&forecast.Temperature,
			&forecast.Date,
			&forecast.DetailInfo,
		)
		if err != nil {
			return nil, fmt.Errorf("repository.Forecasts.GetAll: %w", err)
		}
		forecasts = append(forecasts, forecast)
	}
	return forecasts, nil
}
