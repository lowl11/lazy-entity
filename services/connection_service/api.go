package connection_service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"log"
	"time"
)

func (service *Service) ConnectionPool() (*sqlx.DB, error) {
	if service.connectionPool != nil {
		return service.connectionPool, nil
	}

	// connection pool for Postgres
	connectionPool, err := sql_helper.ConnectionPool(
		service.connectionString,
		service.maxConnections,
		service.maxLifetime,
	)
	if err != nil {
		return nil, err
	}

	// ping database
	log.Println("Ping Postgres database connection pool...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = connectionPool.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Println("Ping Postgres database connection pool done!")

	service.connectionPool = connectionPool
	return connectionPool, nil
}

func (service *Service) Connection() (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// connection for Postgres
	connection, err := sql_helper.Connection(
		service.connectionString,
		service.maxConnections,
		service.maxLifetime,
		time.Second*5,
	)
	if err != nil {
		return nil, err
	}

	log.Println("Ping Postgres database connection...")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = connection.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Println("Ping Postgres database connection done!")

	return connection, nil
}

func (service *Service) SetMaxConnections(maxConnections int) *Service {
	service.maxConnections = maxConnections
	return service
}

func (service *Service) SetMaxLifetime(maxLifetime int) *Service {
	service.maxLifetime = maxLifetime
	return service
}
