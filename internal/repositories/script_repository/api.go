package script_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

func (repo *Repository) Script(folder, script string) string {
	return repo.script.GetScript(folder, script)
}

func (repo *Repository) StartScript(script string) string {
	return repo.script.GetStartScript(script)
}

func (repo *Repository) Guid() string {
	return repo.base.Guid()
}

func (repo *Repository) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	return repo.base.Ctx(customTimeout...)
}

func (repo *Repository) CloseRows(rows *sqlx.Rows) {
	repo.base.CloseRows(rows)
}

func (repo *Repository) Rollback(transaction *sqlx.Tx) {
	repo.base.Rollback(transaction)
}

func (repo *Repository) Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	return repo.base.Transaction(connection, transactionActions)
}
