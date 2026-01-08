package initialize

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/anhnv1202/base-go/global"
)

type cleanupFunc func()

var cleanups []cleanupFunc

func OnShutdown(fn cleanupFunc) {
	cleanups = append(cleanups, fn)
}

func Shutdown() {
	for i := len(cleanups) - 1; i >= 0; i-- {
		cleanups[i]()
	}
}

func WaitForSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("Shutting down...")
}
