package zap_logger

import (
	"github.com/wrlin1218/url_shortener/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

/**
日志初始化
*/

func Init(option logger.LogOption) {
	logger.LogInstance = NewZapLogger(option)
	logger.LogInstance.Info("Zap logger init success.")
}

func NewZapLogger(option logger.LogOption) *ZapLogger {
	// 1. parse option
	var cores []zapcore.Core
	switch option.Writter {
	case "file":
		cores = append(cores, FileCore(option))
	case "console":
		cores = append(cores, ConsoleCore(option))
	default:
		cores = append(cores, FileCore(option))
		cores = append(cores, ConsoleCore(option))
	}

	// 2. combine and create logger
	core := zapcore.NewTee(cores...)
	return &ZapLogger{logger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))}
}

/**
不同等级的日志入口方法实现
*/

func (z *ZapLogger) Log(level logger.LogLevel, msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			zapFields = append(zapFields, zap.Any(fields[i].(string), fields[i+1]))
		}
	}

	switch level {
	case logger.DebugLevel:
		z.logger.Debug(msg, zapFields...)
	case logger.InfoLevel:
		z.logger.Info(msg, zapFields...)
	case logger.WarnLevel:
		z.logger.Warn(msg, zapFields...)
	case logger.ErrorLevel:
		z.logger.Error(msg, zapFields...)
	case logger.FatalLevel:
		z.logger.Fatal(msg, zapFields...)
	default:
		z.logger.Info(msg, zapFields...)
	}
}

func (z *ZapLogger) Debug(msg string, fields ...interface{}) {
	z.Log(logger.DebugLevel, msg, fields...)
}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	z.Log(logger.InfoLevel, msg, fields...)
}

func (z *ZapLogger) Warn(msg string, fields ...interface{}) {
	z.Log(logger.WarnLevel, msg, fields...)
}

func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	z.Log(logger.ErrorLevel, msg, fields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...interface{}) {
	z.Log(logger.FatalLevel, msg, fields...)
}
