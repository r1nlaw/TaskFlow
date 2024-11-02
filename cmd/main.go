package main

import (
	taskflow "TaskFlow"
	"TaskFlow/package/handler"
	"TaskFlow/package/repository"
	"TaskFlow/package/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(taskflow.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http setver: %s", err.Error())
	}

}
