package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syougo1209/ulife-share/controller"
	"github.com/syougo1209/ulife-share/model/repository"
)

var ur = repository.NewUserRepository()
var uc = controller.NewUserController(ur)

func main() {
	router := gin.Default()
	router.POST("/users", uc.PostUser)
	router.GET("/users", uc.ReturnError)
	router.Run(":8080")
}
