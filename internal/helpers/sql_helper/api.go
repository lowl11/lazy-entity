package sql_helper

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func ContainsAggregate(value string) bool {
	upper := strings.ToUpper(value)

	return strings.Contains(upper, aggregateCount) ||
		strings.Contains(upper, aggregateMin) ||
		strings.Contains(upper, aggregateMax)
}

func ConnectionPool(connectionString string, maxConnections, maxLifetime int) (*sqlx.DB, error) {
	// connection pool for Postgres
	connectionPool, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection pool configurations
	connectionPool.SetMaxOpenConns(maxConnections)
	connectionPool.SetMaxIdleConns(maxConnections)
	connectionPool.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Minute)

	return connectionPool, nil
}

func Connection(connectionString string, maxConnections, maxLifetime int, timeout time.Duration) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// connection for Postgres
	connection, err := sqlx.ConnectContext(ctx, "postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection configurations
	connection.SetMaxOpenConns(maxConnections)
	connection.SetMaxIdleConns(maxConnections)
	connection.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Minute)

	return connection, nil
}

func CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Println(err)
	}
}

func Rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(err.Error(), "sql: transaction has already been committed or rolled back") {
			log.Println(err, "Rollback transaction error")
		}
	}
}

func Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := connection.Beginx()
	if err != nil {
		return err
	}
	defer Rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	defaultTimeout := time.Second * 5
	if len(customTimeout) > 0 {
		defaultTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), defaultTimeout)
}
