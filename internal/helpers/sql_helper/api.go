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

func IsKeyword(word string) bool {
	for _, keyword := range keywords {
		if keyword == word {
			return true
		}
	}

	return false
}

func AliasName(name string) string {
	buildAlias := func(before, after string) string {
		alias := strings.Builder{}
		alias.Grow(len(before) + len(after) + 10)
		alias.WriteString("\"")
		alias.WriteString(before)
		alias.WriteString("\".")
		alias.WriteString(after)
		return alias.String()
	}

	if strings.Contains(name, ".") {
		before, after, _ := strings.Cut(name, ".")
		if IsKeyword(before) {
			return buildAlias(before, after)
		}
	}

	if IsKeyword(name) {
		if strings.Contains(name, ".") {
			before, after, _ := strings.Cut(name, ".")
			return buildAlias(before, after)
		}

		alias := strings.Builder{}
		alias.Grow(len(name) + 4)
		alias.WriteString("\"")
		alias.WriteString(name)
		alias.WriteString("\"")
		return alias.String()
	}

	return name
}

func ConditionAlias(aliasName, conditions string) string {
	search := aliasName + "."
	return strings.ReplaceAll(conditions, search, AliasName(aliasName)+".")
}

func Rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(
			err.Error(),
			"sql: transaction has already been committed or rolled back",
		) {
			log.Println(err, "Rollback transaction error")
		}
	}
}
