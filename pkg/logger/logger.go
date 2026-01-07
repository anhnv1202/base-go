package logger

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/anhnv1202/base-go/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const dateFormat = "02-01-2006"
const timeFormat = "02/01/2006 15:04:05.000 -07:00"

var whitespaceRegex = regexp.MustCompile(`[\t\n\r]+`)

func sanitizeString(s string) string {
	return strings.TrimSpace(whitespaceRegex.ReplaceAllString(s, " "))
}

type Logger struct {
	*zap.Logger
	rotators []*DailyRotateLogger
}

type DailyRotateLogger struct {
	*lumberjack.Logger
	baseDir     string
	currentDate string
	mu          sync.Mutex
}

func NewLogger(config setting.LogConfig) *Logger {
	level := getLogLevel(config.Level)
	encoder := newEncoderConfig()

	// Console core - format based on config
	var consoleEncoder zapcore.Encoder
	if config.Format == "json" {
		consoleEncoder = zapcore.NewJSONEncoder(encoder)
	} else {
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder.ConsoleSeparator = " | "
		consoleEncoder = zapcore.NewConsoleEncoder(encoder)
	}

	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), level),
	}

	// File cores - always JSON, separate by level
	rotators := []*DailyRotateLogger{}
	jsonEncoder := zapcore.NewJSONEncoder(newEncoderConfig())

	for _, l := range []struct {
		lvl    zapcore.Level
		name   string
	}{
		{zapcore.DebugLevel, "debug"},
		{zapcore.InfoLevel, "info"},
		{zapcore.WarnLevel, "warn"},
		{zapcore.ErrorLevel, "error"},
		{zapcore.FatalLevel, "fatal"},
	} {
		if level <= l.lvl {
			writer := newDailyRotateLogger(config, l.name)
			rotators = append(rotators, writer)
			cores = append(cores, zapcore.NewCore(jsonEncoder, zapcore.AddSync(writer), l.lvl))
		}
	}

	return &Logger{
		Logger:   zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)),
		rotators: rotators,
	}
}

func newDailyRotateLogger(config setting.LogConfig, level string) *DailyRotateLogger {
	today := time.Now().Format(dateFormat)
	baseDir := filepath.Join("logs", level)
	os.MkdirAll(baseDir, 0755)

	return &DailyRotateLogger{
		baseDir:     baseDir,
		currentDate: today,
		Logger: &lumberjack.Logger{
			Filename:   filepath.Join(baseDir, today+".log"),
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
			LocalTime:  true,
		},
	}
}

func (d *DailyRotateLogger) Write(p []byte) (n int, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	today := time.Now().Format(dateFormat)
	if d.currentDate != today {
		d.Logger.Close()
		d.currentDate = today
		d.Logger.Filename = filepath.Join(d.baseDir, today+".log")
		d.Logger.Rotate()
	}

	return d.Logger.Write(p)
}

func (d *DailyRotateLogger) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.Logger.Close()
}

func newEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	cfg.TimeKey = "time"
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	return cfg
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *Logger) Sync() error {
	return l.Logger.Sync()
}

func (l *Logger) Close() error {
	l.Sync()
	for _, r := range l.rotators {
		r.Close()
	}
	return nil
}

func (l *Logger) with(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...), rotators: l.rotators}
}

func (l *Logger) WithRequestID(id string) *Logger {
	return l.with(zap.String("request_id", id))
}

func (l *Logger) WithUserID(id string) *Logger {
	return l.with(zap.String("user_id", id))
}

func (l *Logger) WithError(err error) *Logger {
	if err == nil {
		return l
	}
	return l.with(zap.String("error", sanitizeString(err.Error())))
}

func (l *Logger) Err(msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.String("error", sanitizeString(err.Error())))
	}
	l.Logger.Error(msg, fields...)
}

func (l *Logger) WithContext(fields map[string]any) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return l.with(zapFields...)
}

func (l *Logger) LogDBQuery(query string, duration time.Duration, rows int64) {
	l.Debug("DB Query",
		zap.String("query", query),
		zap.Duration("duration", duration),
		zap.Int64("rows", rows),
	)
}
