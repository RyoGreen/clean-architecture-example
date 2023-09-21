package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBConfig DBconfig `json:"db"`
}

type DBconfig struct {
	User     string `json:"user"`
	DBName   string `json:"name"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func NewConfig() (*Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
