package main

import (
	"net/http"

	"github.com/syougo1209/ulife-share/controller"
	"github.com/syougo1209/ulife-share/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

var ur = repository.NewUserRepository()
var uc = controller.NewUsersController(ur)
var user_router = controller.NewRouter(uc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/users/", user_router.HandleUsersRequest)
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
