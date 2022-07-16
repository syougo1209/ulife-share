package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syougo1209/ulife-share/controller/dto"
	"github.com/syougo1209/ulife-share/model/entity"
	"github.com/syougo1209/ulife-share/model/repository"
)

type UserController interface {
	PostUser(c *gin.Context)
}

type userController struct {
	ur repository.UserRepository
}

func NewUserController(ur repository.UserRepository) UserController {
	return &userController{ur}
}

func (uc *userController) PostUser(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.BindJSON(&userRequest); err != nil {
		return
	}
	user := entity.UserEntity{Name: userRequest.Name, Email: userRequest.Email, PasswordHash: userRequest.Password, Introduction: userRequest.Introduction}
	if err := uc.ur.InsertUser(user); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}
