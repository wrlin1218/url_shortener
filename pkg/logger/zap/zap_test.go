package zap_logger

import (
	"testing"

	"github.com/wrlin1218/url_shortener/pkg/logger"
)

func TestInit(t *testing.T) {
	Init(logger.LogOption{
		Writter:    "all",
		Filename:   "test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	})
	logger.LogInstance.Debug("Test")
}
