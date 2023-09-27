package main

import (
	"clean-architecture/config"
	"clean-architecture/controller"
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
	userRepo := repo.NewUserRepo(db)
	userUseCase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUseCase)
	postRepo := repo.NewPostRepo(db)
	postUseCase := usecase.NewPostUsecase(postRepo)
	postController := controller.NewPostController(postUseCase)
	e := router.NewRouter(userController, postController)
	e.Logger.Error(e.Start(":8080"))
}
