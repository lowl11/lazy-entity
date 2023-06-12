package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/crud_repository"
)

type ICrudRepository[T any, ID repositories.IComparableID] interface {
	IUniversalRepository[T, ID]

	GetAll() ([]T, error)
	GetByID(id ID) (*T, error)
	GetByIdList(id []ID) ([]T, error)

	UpdateByID(id ID, entity T) error

	DeleteAll() error
	DeleteByID(id ID) error
}

func NewCrud[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string, params ...string) ICrudRepository[T, ID] {
	var alias string
	var idName string

	if len(params) > 0 {
		alias = params[0]
	}
	if len(params) > 1 {
		idName = params[1]
	}

	newRepo := crud_repository.New[T, ID](connection, tableName)
	newRepo.Alias(alias).IdName(idName)
	return newRepo
}
