package logger

/**
通用日志接口定义
*/

type LogOption struct {
	Writter    string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type Logger interface {
	Log(level LogLevel, msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

/**
全局日志对象 + 方法
*/

var LogInstance Logger

// 全局方法实现
func Log(level LogLevel, msg string, fields ...interface{}) {
	LogInstance.Log(level, msg, fields...)
}

func Debug(msg string, fields ...interface{}) {
	LogInstance.Debug(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	LogInstance.Info(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	LogInstance.Warn(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	LogInstance.Error(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	LogInstance.Fatal(msg, fields...)
}
