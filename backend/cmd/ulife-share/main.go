package main

import (
	"github.com/syougo1209/ulife-share/controller"
	"github.com/syougo1209/ulife-share/model/repository"
)

var ur = repository.NewUserRepository()
var uc = controller.NewUserController(ur)

func main() {
	router.POST("/users", uc.PostUser)
	router.Run(":8080")
}
