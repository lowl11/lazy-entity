package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
)

// IRepository is base repository for any SQL based repository.
// It includes basic methods for writing any CRUD method
type IRepository interface {
	// Guid returns generated new UUID.
	// Usually, it uses for generating some new entity IDs.
	Guid() string

	// Ctx returns context with timeout.
	// Argument is an array of durations, but takes only first element.
	// By default it is 30 seconds
	Ctx(customTimeout ...time.Duration) (context.Context, func())

	// CloseRows prints closing rows error
	CloseRows(rows *sqlx.Rows)

	// Rollback do transaction ROLLBACK.
	// It is not error if it was already commited
	Rollback(transaction *sqlx.Tx)

	// TransactionQuery provides transaction (*sqlx.Tx) object to transactionActions.
	// COMMIT and ROLLBACK calls automatically
	TransactionQuery(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error
}

func NewBase() IRepository {
	return &base_repository.Repository{}
}
