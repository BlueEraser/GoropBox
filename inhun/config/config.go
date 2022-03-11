package config

import (
	"encoding/json"
	"io/ioutil"
)

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Port     int    `json:"port"`
	TimeZone string `json:"timezone"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

func LoadConfig(filepath string) (*Config, error) {
	cfg := &Config{}

	dataBytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		return cfg, err
	}

	json.Unmarshal(dataBytes, cfg)

	return cfg, nil
}
