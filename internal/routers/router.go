package routers

import (
	c "github.com/anhnv1202/base-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	v1.GET("/ping", c.NewUserController().GetUser)
	return router
}

