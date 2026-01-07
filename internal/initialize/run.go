package initialize

import (
	"strconv"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/internal/routers"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	defer global.Logger.Close() // Ensure cleanup on shutdown

	InitDatabase()
	InitRedis()

	r := routers.InitRouter()
	global.Logger.Info("starting server", zap.Int("port", global.Config.App.Port))

	if err := r.Run(":" + strconv.Itoa(global.Config.App.Port)); err != nil {
		global.Logger.Error("server stopped", zap.Error(err))
	}
}
