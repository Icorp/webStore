package main

import (
	"log"

	webstore "github.com/Icorp/webStore"
	handler "github.com/Icorp/webStore/pkg/handlers"
	"github.com/Icorp/webStore/pkg/repository"
	"github.com/Icorp/webStore/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(webstore.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
