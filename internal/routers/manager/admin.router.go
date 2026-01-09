package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ur *AdminRouter) RegisterAdminRouter(router *gin.RouterGroup) {
	//pulic router
	adminRouterPublic := router.Group("/admins")
	{
		adminRouterPublic.POST("/login")
	}

	adminRouterPrivate := router.Group("/admin/user")
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
