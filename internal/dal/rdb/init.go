package rdb

import (
	"github.com/glebarez/sqlite"
	"github.com/wrlin1218/url_shortener/internal/models"
	"github.com/wrlin1218/url_shortener/pkg/logger"
	"gorm.io/gorm"
)

var rdb *gorm.DB

func GetRDB() *gorm.DB {
	return rdb
}

/**
gorm本身对sql操作进行了较为完善的封装，因此此处仅处理初始化即可
*/

type RDBOption struct {
	Dialact string // 目前仅支持sqllite
	DSN     string
}

func Init(option RDBOption) *gorm.DB {
	switch option.Dialact {
	case "sqlite":
		// 1. connect
		db, err := gorm.Open(sqlite.Open(option.DSN), &gorm.Config{})
		if err != nil {
			logger.Fatal("Failed to connect sqlite db: %v", err)
		}

		// 2. migrate
		if err := db.AutoMigrate(&models.User{}, &models.Link{}); err != nil {
			logger.Fatal("Failed to migrate database: %v", err)
		}
		rdb = db
	default:
		logger.Fatal("Only support sqlite now")
	}
	return rdb
}
