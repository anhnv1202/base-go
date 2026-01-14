package controllers

import (
	"net/http"

	"github.com/anhnv1202/base-go/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"user": uc.service.GetOne()})
}
