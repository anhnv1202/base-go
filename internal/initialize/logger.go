package initialize

import (
	"github.com/anhnv1202/base-go/global"
	"github.com/anhnv1202/base-go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
