package main

import (
	"net/http"

	"github.com/syougo1209/ulife-share/controller"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/lifes/", controller.NewLifeHandler)

	server.ListenAndServe()
}
