package database_service

import "github.com/jmoiron/sqlx"

type Service struct {
	connectionString string

	maxConnections int
	maxLifetime    int

	connectionPool *sqlx.DB
}

func New(connectionString string) *Service {
	return &Service{
		connectionString: connectionString,

		maxConnections: maxConnections,
		maxLifetime:    maxLifetime,
	}
}
