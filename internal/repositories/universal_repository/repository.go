package universal_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/internal/join_field"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
	"github.com/lowl11/lazy-entity/repo_config"
)

type Repository[T any, ID repositories.IComparableID] struct {
	base_repository.Repository
	connection    *sqlx.DB
	tableName     string
	fieldList     []string
	joinFieldList []join_field.Field

	aliasName string
	idName    string
	joinList  []repo_config.Join

	debug bool
}

func New[T any, ID repositories.IComparableID](
	connection *sqlx.DB,
	tableName string,
) Repository[T, ID] {
	fieldList, joinFieldList := type_helper.GetStructFields[T]()

	return Repository[T, ID]{
		Repository: base_repository.Repository{},

		connection:    connection,
		tableName:     tableName,
		idName:        defaultIdName,
		fieldList:     fieldList,
		joinFieldList: joinFieldList,
		joinList:      make([]repo_config.Join, 0),
	}
}
