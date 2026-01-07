package routers

import (
	c "github.com/anhnv1202/base-go/internal/controllers"
	"github.com/anhnv1202/base-go/internal/middlewares"
	"github.com/anhnv1202/base-go/global"
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
	}

	router := gin.New()
	router.Use(
		middlewares.Recovery(),
		middlewares.Logger(),
		middlewares.AuthMiddleware(),
	)

	v1 := router.Group("/api")
	v1.GET("/ping", c.NewUserController().GetUser)

	return router
}
