package initialize

import (
	"strconv"

	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/internal/routers"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDatabase()
	InitRedis()

	r := routers.InitRouter()
	r.Run(":" + strconv.Itoa(global.Config.App.Port))
}
