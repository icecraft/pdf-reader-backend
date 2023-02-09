package dao

import (
	"context"

	"gorm.io/gorm"
)

type TransactionKey string

func WithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TransactionKey("sinan-transaction"), tx)
}

func GetTransaction(ctx context.Context) *gorm.DB {
	v := ctx.Value(TransactionKey("sinan-transaction"))
	if v == nil {
		return nil
	}
	return v.(*gorm.DB)
}

func RunWithTransaction(ctx context.Context, f func(ctx context.Context, tx *gorm.DB) error) error {
	tx := GetTransaction(ctx)
	if tx == nil {
		return Db.Transaction(func(tx *gorm.DB) error {
			newCtx := WithTransaction(ctx, tx)
			return f(newCtx, tx)
		})
	} else {
		return f(ctx, tx)
	}
}
