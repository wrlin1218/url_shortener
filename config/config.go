// config/config.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
	redis "github.com/wrlin1218/url_shortener/internal/dal/kv/impl"
	"github.com/wrlin1218/url_shortener/internal/dal/rdb"
	"github.com/wrlin1218/url_shortener/pkg/logger"
)

// Config holds the entire configuration
type Config struct {
	Log      logger.LogOption       `mapstructure:"log"`
	Database rdb.RDBOption          `mapstructure:"database"`
	Redis    redis.RedisInitOptions `mapstructure:"redis"`
	Sync     SyncConfig             `mapstructure:"sync"`
}

// SyncConfig holds synchronization configuration
type SyncConfig struct {
	IntervalSeconds int `mapstructure:"interval_seconds"`
}

// LoadConfig reads configuration from file using Viper
func LoadConfig(path, file string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")

	// 设置默认值
	viper.SetDefault("log.writter", "all")
	viper.SetDefault("log.fileName", "logs/app.log")
	viper.SetDefault("log.maxSize", 10)
	viper.SetDefault("log.maxBackup", 5)
	viper.SetDefault("log.maxAge", 30)
	viper.SetDefault("log.compress", false)
	viper.SetDefault("database.dialect", "sqlite")
	viper.SetDefault("database.dsn", "shortlink.db")
	viper.SetDefault("redis.address", "localhost:6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("sync.interval_seconds", 60)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}
