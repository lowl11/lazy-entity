package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/builders/select_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/universal_repository"
)

type IUniversalRepository[T any, ID repositories.IComparableID] interface {
	IRepository

	GetList(conditionFunc func(builder *select_builder.Builder) string, args ...any) ([]T, error)
}

func NewUniversal[T any, ID repositories.IComparableID](
	connection *sqlx.DB,
	tableName string,
	params ...string,
) IUniversalRepository[T, ID] {
	var aliasName string
	var idName string

	if len(params) > 0 {
		aliasName = params[0]
	}
	if len(params) > 1 {
		idName = params[1]
	}

	return universal_repository.New[T, ID](connection, tableName).Alias(aliasName).IdName(idName)
}
