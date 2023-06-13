package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
	"time"
)

type IRepository interface {
	Guid() string
	Ctx(customTimeout ...time.Duration) (context.Context, func())
	CloseRows(rows *sqlx.Rows)
	Rollback(transaction *sqlx.Tx)
	Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error
}

func NewBase() IRepository {
	return &base_repository.Repository{}
}
