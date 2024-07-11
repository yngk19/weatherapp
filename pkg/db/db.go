package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yngk19/weatherapp/internal/config"
)

func ConnectDB(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("ConnectDB config parse: %w", err)
	}
	config.ConnConfig.Host = cfg.Host
	config.ConnConfig.Port = cfg.Port
	config.ConnConfig.Database = cfg.DBName
	config.ConnConfig.User = cfg.User
	config.ConnConfig.Password = cfg.Password
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("ConnectDB: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ConnectDB ping: %w", err)
	}
	return pool, nil
}
