package repo

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const (
	TxKey contextKey = "tx"
)

type BaseRepo interface {
	DB(ctx context.Context) *gorm.DB
}
