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

func (r *Repo) GetByCityID(ctx context.Context, id int) ([]domain.WeatherForecast, error) {
	var forecasts []domain.WeatherForecast
	query := `
		SELECT f.id, f.temp, f.predict_date, f.detail_info 
		FROM forecasts f WHERE f.city_id = $1
		ORDER BY predict_date DESC
		LIMIT 5;
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("repository.Forecasts.GetByCityID: %w", err)
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
			return nil, fmt.Errorf("repository.Forecasts.GetByCityID: %w", err)
		}
		forecasts = append(forecasts, forecast)
	}
	return forecasts, nil
}

func (r *Repo) GetByDate(ctx context.Context, date string, id int) (*domain.WeatherForecast, error) {
	var forecast domain.WeatherForecast
	query := `
		SELECT f.id, f.temp, f.predict_date, f.detail_info
		FROM forecasts f WHERE f.predict_date = $1
		AND f.city_id = $2;
	`
	if err := r.pool.QueryRow(ctx, query, date, id).Scan(&forecast.ID, &forecast.Temperature, &forecast.Date, &forecast.DetailInfo); err != nil {
		return nil, fmt.Errorf("repository.Forecasts.GetByDate: %w", err)
	}
	return &forecast, nil
}
