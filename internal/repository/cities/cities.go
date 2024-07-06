package citiesrepo

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *Repo) Create(ctx context.Context, city dto.Town) error {
	cityId := 0
	query := `
		INSERT INTO cities (name, country, lat, lon)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (name) DO NOTHING
		RETURNING id;
	`
	if err := r.pool.QueryRow(ctx, query, city.Name, city.Country, city.Lat, city.Lon).Scan(&cityId); err != nil {
		return fmt.Errorf("repository.Cities.Create: %w", err)
	}
	return nil
}
