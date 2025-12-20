package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBConfig DBconfig `json:"db"`
	HashKey  string   `json:"hash_key"`
	BlockKey string   `json:"block_key"`
}

type DBconfig struct {
	User     string `json:"user"`
	DBName   string `json:"dbname"`
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
