package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type IRepository interface {
	Guid() string
	Ctx(customTimeout ...time.Duration) (context.Context, func())
	CloseRows(rows *sqlx.Rows)
	Rollback(transaction *sqlx.Tx)
	Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error
}
