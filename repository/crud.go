package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/crud_repository"
)

type ICrudRepository[T any, ID repositories.IComparableID] interface {
	IRepository

	Count() (int, error)
	ExistByID(id ID) (bool, error)

	GetAll() ([]T, error)
	GetByID(id ID) (*T, error)

	Add(entity any) (ID, error)
	AddList(entityList []any) error

	SaveByID(id ID, entity T) error

	SaveByCondition(
		conditionFunc func(builder *update_builder.Builder) string,
		entity T,
	) error

	UpdateByID(id ID, updateEntity any) error
	UpdateByCondition(
		conditionFunc func(builder *update_builder.Builder) string,
		updateEntity any,
	) error

	DeleteAll() error
	DeleteByID(id ID) error
}

func NewCrud[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string, params ...string) *crud_repository.CrudRepository[T, ID] {
	var alias string
	var idName string

	if len(params) > 0 {
		alias = params[0]
	}
	if len(params) > 1 {
		idName = params[1]
	}
	return crud_repository.New[T, ID](connection, tableName).Alias(alias).IdName(idName)
}
