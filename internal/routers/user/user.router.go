package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) RegisterUserRouter(router *gin.RouterGroup) {
	//pulic router
	userRouterPublic := router.Group("/users")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/login")
	}

	userRouterPrivate := router.Group("/users")
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/profile")
	}
}
