package base_repository

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (repo *Repository) Guid() string {
	return uuid.New().String()
}

func (repo *Repository) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	defaultTimeout := time.Second * 30
	if len(customTimeout) > 0 {
		defaultTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), defaultTimeout)
}

func (repo *Repository) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Println(err)
	}
}

func (repo *Repository) Rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(
			err.Error(),
			"sql: transaction has already been committed or rolled back",
		) {
			log.Println(err, "Rollback transaction error")
		}
	}
}

func (repo *Repository) Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := connection.Beginx()
	if err != nil {
		return err
	}
	defer repo.Rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
