package citiesrepo

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

func (r *Repo) Create(ctx context.Context, city dto.Town) error {
	var id int
	query := `
		INSERT INTO cities (name, country, state, lat, lon)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (lat, lon) DO NOTHING
		RETURNING id;
	`
	if err := r.pool.QueryRow(ctx, query, city.Name, city.Country, city.State, city.Lat, city.Lon).Scan(&id); err != nil {
		return fmt.Errorf("repository.Cities.Create: %w", err)
	}
	return nil
}

func (r *Repo) GetAll(ctx context.Context) ([]domain.Town, error) {
	var towns []domain.Town
	query := `
		SELECT c.id, c.name, c.country, c.state, c.lat, c.lon  
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
			&town.State,
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
