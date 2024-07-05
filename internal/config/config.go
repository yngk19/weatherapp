package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	DBConfig
	OpenWeatherAPIKey string
}

type ServerConfig struct {
	Env        string `yaml:"env" env-default:"prod"`
	HTTPServer `yaml:"http_server"`
}

type DBConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	SSLMode        string
	DBName         string
	MigrationsPath string
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
	Port         int           `yaml:"port" env-default:"80"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load the .env file!: %s", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	OpenWeatherAPIKey := os.Getenv("OPENWEATHER_KEY")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbSSLMode := os.Getenv("SSL_MODE")
	dbMigrationsPath := os.Getenv("MIGRATIONS_PATH")
	var db DBConfig = DBConfig{
		Host:           dbHost,
		Port:           dbPort,
		User:           dbUser,
		Password:       dbPassword,
		SSLMode:        dbSSLMode,
		DBName:         dbName,
		MigrationsPath: dbMigrationsPath,
	}
	if configPath == "" {
		log.Fatal("Config file is not set!")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist!: %s", err)
	}
	var srv ServerConfig
	if err := cleanenv.ReadConfig(configPath, &srv); err != nil {
		log.Fatalf("Cannot read the config!: %s", err)
	}
	var cfg Config = Config{
		DBConfig:          db,
		ServerConfig:      srv,
		OpenWeatherAPIKey: OpenWeatherAPIKey,
	}
	return &cfg
}
