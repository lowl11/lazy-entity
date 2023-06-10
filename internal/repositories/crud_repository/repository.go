package crud_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
)

type CrudRepository[T any, ID repositories.IComparableID] struct {
	base_repository.Repository
	connection *sqlx.DB
	tableName  string
	aliasName  string
	idName     string
	fieldList  []string
}

func New[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string) *CrudRepository[T, ID] {
	return &CrudRepository[T, ID]{
		connection: connection,
		tableName:  tableName,
		idName:     defaultIdName,
		fieldList:  type_helper.GetStructFields[T](),
	}
}
