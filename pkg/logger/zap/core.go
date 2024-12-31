package zap_logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/wrlin1218/url_shortener/pkg/logger"
	"go.uber.org/zap/zapcore"
)

/**
对core进行初始化
*/

func ConsoleCore(option logger.LogOption) zapcore.Core {
	// 1. init writter
	consoleWriter := zapcore.AddSync(os.Stdout)

	// 2. config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 日志级别大写（INFO, DEBUG）
		EncodeTime:   zapcore.ISO8601TimeEncoder,  // 时间格式 ISO8601
		EncodeCaller: customCallerEncoder,         // [文件名|函数名|行号]
	}

	// 3. create zapcore
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 控制台日志使用 Console 格式
		consoleWriter,                            // 输出到控制台
		zapcore.DebugLevel,                       // 最低日志级别
	)
}

func FileCore(option logger.LogOption) zapcore.Core {
	// 1. init writter
	fileLogger := &lumberjack.Logger{
		Filename:   option.Filename,
		MaxSize:    option.MaxSize,
		MaxBackups: option.MaxBackups,
		MaxAge:     option.MaxAge,
		Compress:   option.Compress,
	}
	fileWritter := zapcore.AddSync(fileLogger)

	// 2. config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 日志级别大写（INFO, DEBUG）
		EncodeTime:   zapcore.ISO8601TimeEncoder,  // 时间格式 ISO8601
		EncodeCaller: customCallerEncoder,         // [文件名|函数名|行号]
	}

	// 3. create zapcore
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		fileWritter,
		zapcore.DebugLevel,
	)
}

// customCallerEncoder 自定义encode，调整日志输出格式，并加上方法信息
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// 获取函数名
	pc := caller.PC
	funcName := "unknown"
	if f := runtime.FuncForPC(pc); f != nil {
		funcName = f.Name()
		if idx := strings.LastIndex(f.Name(), "."); idx != -1 {
			funcName = f.Name()[idx+1:]
		} else if idx := strings.LastIndex(funcName, "/"); idx != -1 {
			funcName = f.Name()[idx+1:]
		}
	}

	// 格式化 caller 为 "[文件名|函数名|行号]"
	enc.AppendString(fmt.Sprintf("[%s|%s|%d]", TrimmedPath(caller), funcName, caller.Line))
}

// TrimmedPath returns a file description of the caller,
func TrimmedPath(ec zapcore.EntryCaller) string {
	if !ec.Defined {
		return "undefined"
	}
	// 文件名
	idx := strings.LastIndexByte(ec.File, '/')
	if idx == -1 {
		return ec.FullPath()
	}
	// 目录名
	idx = strings.LastIndexByte(ec.File[:idx], '/')
	if idx == -1 {
		return ec.FullPath()
	}
	return ec.File[idx+1:]
}
