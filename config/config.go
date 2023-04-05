package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"9090"`
	Database string `env:"DATABASE" envDefault:"onelab"`
	PgURL    string `env:"PG_URL" envDefault:"user=user password=secret dbname=db sslmode=disable host=localhost port=5432"`
}

//где файлик ? 
func New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err) // можно вернуть ошибку либо убрать возврат ошибки
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err) // можно вернуть ошибку либо убрать возврат ошибки
	}
	return &cfg, nil
}
