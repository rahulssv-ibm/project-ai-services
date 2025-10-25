package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger      *zap.Logger
	atomicLevel zap.AtomicLevel
)

func init() {
	atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "",
		LevelKey:       "",
		NameKey:        "",
		CallerKey:      "",
		MessageKey:     "msg",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	})

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), atomicLevel)
	Logger = zap.New(core)
}

func GetLogger() *zap.Logger {
	return Logger
}

func SetLogLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func GetLogLevel() zapcore.Level {
	return atomicLevel.Level()
}
