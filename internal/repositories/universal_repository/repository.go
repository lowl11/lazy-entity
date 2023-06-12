package universal_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
)

type Repository[T any, ID repositories.IComparableID] struct {
	base_repository.Repository
	connection *sqlx.DB
	tableName  string
	fieldList  []string

	aliasName string
	idName    string
}

func New[T any, ID repositories.IComparableID](
	connection *sqlx.DB,
	tableName string,
) Repository[T, ID] {
	return Repository[T, ID]{
		connection: connection,
		tableName:  tableName,
		idName:     defaultIdName,
		fieldList:  type_helper.GetStructFields[T](),
	}
}
