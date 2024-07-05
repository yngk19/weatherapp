package main

import (
	"fmt"

	"github.com/yngk19/weatherapp/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.OpenWeatherAPIKey)

}
