package db

import (
	"clean-architecture/config"
	"database/sql"
	"fmt"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.DBConfig.User, cfg.DBConfig.DBName, cfg.DBConfig.Password, cfg.DBConfig.Host, cfg.DBConfig.Port))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	userQuery := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);`
	postQuery := `CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    user_id INT REFERENCES users(id)
);`

	if _, err := db.Exec(userQuery); err != nil {
		return nil, err
	}
	if _, err := db.Exec(postQuery); err != nil {
		return nil, err
	}
	return db, nil
}
