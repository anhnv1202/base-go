package manager

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) RegisterUserRouter(router *gin.RouterGroup) {
	//pulic router
	// userRouterPublic := router.Group("/managers")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/login")
	// }

	userRouterPrivate := router.Group("/admin/user")
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.POST("/active")
	}
}
