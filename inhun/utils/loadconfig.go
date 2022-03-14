package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Url string `json:"url"`
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
