package controllers

import (
	"net/http"

	"github.com/anhnv1202/base-go/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"user": uc.service.GetOne()})
}
