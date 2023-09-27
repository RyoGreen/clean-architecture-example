package main

import (
	"clean-architecture/config"
	"clean-architecture/db"
	"clean-architecture/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logger.L.Error(err.Error())
		return
	}
	db, err := db.NewDB(cfg)
	if err != nil {
		logger.L.Error(err.Error())
		return
	}
	defer db.Close()

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
		logger.L.Error(err.Error())
		return
	}
	if _, err := db.Exec(postQuery); err != nil {
		logger.L.Error(err.Error())
		return
	}
}
