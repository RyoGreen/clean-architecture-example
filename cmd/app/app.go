package main

import (
	"clean-architecture/config"
	"clean-architecture/controller"
	"clean-architecture/cookie"
	"clean-architecture/db"
	"clean-architecture/logger"
	"clean-architecture/repo"
	"clean-architecture/router"
	"clean-architecture/usecase"
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
	cm := cookie.NewCookieManager(cfg)
	userRepo := repo.NewUserRepo(db, cm)
	userUseCase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUseCase)
	postRepo := repo.NewPostRepo(db)
	postUseCase := usecase.NewPostUsecase(postRepo)
	postController := controller.NewPostController(postUseCase)
	sessionRepo := repo.NewSessionRepo(db, cm)
	sessionUseCase := usecase.NewSessionUsecase(sessionRepo)
	sessionController := controller.NewSessionController(sessionUseCase)
	e := router.NewRouter(userController, postController, sessionController)
	e.Logger.Error(e.Start(":8080"))
}
