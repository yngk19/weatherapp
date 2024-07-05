package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yngk19/weatherapp/internal/config"
	"github.com/yngk19/weatherapp/internal/pkg/db"
	citiesrepo "github.com/yngk19/weatherapp/internal/repository/cities"
)

func main() {
	cfg := config.MustLoad()
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB connected!")
	_ = citiesrepo.New(pool)
}
