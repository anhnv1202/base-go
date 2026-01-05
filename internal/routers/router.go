package routers

import (
	c "github.com/anhnv1202/base-go/internal/controllers"
	"github.com/anhnv1202/base-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())
	v1 := router.Group("/api")
	v1.GET("/ping", c.NewUserController().GetUser)
	return router
}

