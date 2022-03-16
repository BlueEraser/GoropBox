package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Database struct {
		HOST     string `env:"DATABASE.HOST"`
		PORT     string `env:"DATABASE.PORT"`
		DBNAME   string `env:"DATABASE.DBNAME"`
		USERNAME string `env:"DATABASE.USERNAME"`
		PASSWORD string `env:"DATABASE.PASSWORD"`
	}
}

var Cfg *Config

func Init() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Load .env file failed")
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", config)
	Cfg = &config
}
