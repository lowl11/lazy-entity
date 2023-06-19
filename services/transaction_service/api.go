package transaction_service

import "github.com/lowl11/lazy-entity/internal/helpers/sql_helper"

func (service *Service) Run() error {
	tx, err := service.connection.Beginx()
	if err != nil {
		return err
	}
	defer sql_helper.Rollback(tx)

	if err = service.actions(tx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
