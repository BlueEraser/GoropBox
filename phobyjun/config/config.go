package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		HOST     string `env:"DATABASE.HOST"`
		PORT     string `env:"DATABASE.PORT"`
		DBNAME   string `env:"DATABASE.DBNAME"`
		USER     string `env:"DATABASE.USER"`
		PASSWORD string `env:"DATABASE.PASSWORD"`
	}
	Session struct {
		SECRET string `env:"SESSION.SECRET"`
	}
}

var Cfg *Config

func init() {
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
