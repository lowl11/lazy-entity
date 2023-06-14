package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories/script_repository"
	"github.com/lowl11/lazy-entity/services/script_service"
	"time"
)

type IScriptRepository interface {
	Script(folder, script string) string
	StartScript(script string) string

	Guid() string
	Ctx(customTimeout ...time.Duration) (context.Context, func())
	CloseRows(rows *sqlx.Rows)
	Rollback(transaction *sqlx.Tx)
	Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error
}

func NewScript(scriptService *script_service.Service) *script_repository.Repository {
	return script_repository.New(scriptService)
}
