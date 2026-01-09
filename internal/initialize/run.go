package initialize

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/anhnv1202/base-go/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDatabase()
	InitRedis()

	r := InitRouter()
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(global.Config.App.Port),
		Handler: r,
	}

	go func() {
		global.Logger.Info("starting server", zap.Int("port", global.Config.App.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("server error", zap.Error(err))
		}
	}()

	WaitForSignal()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("server shutdown error", zap.Error(err))
	}

	Shutdown()
}
