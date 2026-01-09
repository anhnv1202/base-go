package initialize

import (
	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/internal/middlewares"
	"github.com/anhnv1202/base-go/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	ginModes := map[string]string{
		"production": gin.ReleaseMode,
		"test":       gin.TestMode,
	}

	if mode, ok := ginModes[global.Config.App.Env]; ok {
		gin.SetMode(mode)
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	}

	router := gin.New()
	router.Use(
		middlewares.Recovery(),
		middlewares.Logger(),
		middlewares.AuthMiddleware(),
	)

	v1 := router.Group("/api")
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User
	{
		v1.GET("/Health")
	}
	{
		managerRouter.RegisterAdminRouter(v1)
		managerRouter.RegisterUserRouter(v1)
	}
	{
		userRouter.RegisterUserRouter(v1)
	}

	return router
}
