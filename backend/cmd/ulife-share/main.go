package main

import (
	"net/http"

	application "github.com/syougo1209/ulife-share/application/usecase"
	"github.com/syougo1209/ulife-share/controller"
	"github.com/syougo1209/ulife-share/infrastructure/datastore"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	lifeRepo := datastore.NewLifeRepository()
	lifeUC := application.NewLifeUseCase(lifeRepo)
	lifeHandler := controller.NewLifeHandler(lifeUC)

	http.HandleFunc("/lifes/", controller.NewLifeHandler)

	server.ListenAndServe()
}
