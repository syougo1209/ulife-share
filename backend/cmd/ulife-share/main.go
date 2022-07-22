package main

import (
	"net/http"

	application "github.com/syougo1209/ulife-share/application/usecase"
	"github.com/syougo1209/ulife-share/infrastructure/datastore"
	"github.com/syougo1209/ulife-share/interfaces/api/handler"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	lifeRepo := datastore.NewLifeRepository()
	lifeUC := application.NewLifeUseCase(lifeRepo)

	http.HandleFunc("/lifes/", func(w http.ResponseWriter, r *http.Request) {
		handler.NewLifeHandler(w, r, lifeUC)
	})

	server.ListenAndServe()
}
