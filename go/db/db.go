package db

import (
	"clean-architecture/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.DBConfig.User, cfg.DBConfig.DBName, cfg.DBConfig.Password, cfg.DBConfig.Host, cfg.DBConfig.Port))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
