package base_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"time"
)

func (repo *Repository) Guid() string {
	return uuid.New().String()
}

func (repo *Repository) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	return sql_helper.Ctx(customTimeout...)
}

func (repo *Repository) CloseRows(rows *sqlx.Rows) {
	sql_helper.CloseRows(rows)
}

func (repo *Repository) Rollback(transaction *sqlx.Tx) {
	sql_helper.Rollback(transaction)
}

func (repo *Repository) Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	return sql_helper.Transaction(connection, transactionActions)
}
