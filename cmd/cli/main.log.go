package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger := getEncoderLog()
	// logger.Info("Hello, world!")
    logger := getEncoderLog()
    logger.Info("Hello, world!")
    logger.Error("This is an error message")
}

func getEncoderLog() *zap.Logger {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encodeConfig), getWriterSync(), zap.InfoLevel))
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole,syncFile)
}
