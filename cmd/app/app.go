package main

import (
	"clean-architecture/config"
	"clean-architecture/db"
	"clean-architecture/logger"
	"clean-architecture/router"
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
	e := router.NewRouter()
	e.Logger.Error(e.Start(":8080"))
}
