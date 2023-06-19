package transaction_service

import "github.com/jmoiron/sqlx"

type Service struct {
	connection *sqlx.DB
	actions    func(tx *sqlx.Tx) error
}

func New(connection *sqlx.DB, actions func(tx *sqlx.Tx) error) *Service {
	return &Service{
		connection: connection,
		actions:    actions,
	}
}
