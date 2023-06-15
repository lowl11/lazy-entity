package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/crud_repository"
	"github.com/lowl11/lazy-entity/repo_config"
)

type ICrudRepository[T any, ID repositories.IComparableID] interface {
	IUniversalRepository[T, ID]

	ExistByID(id ID) (bool, error)

	GetAll() ([]T, error)
	GetByID(id ID) (*T, error)
	GetByIdList(id []ID) ([]T, error)

	UpdateByID(id ID, entity T) error

	DeleteAll() error
	DeleteByID(id ID) error
}

func NewCrud[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string, config repo_config.Crud) ICrudRepository[T, ID] {
	newRepo := crud_repository.New[T, ID](connection, tableName)
	newRepo.Alias(config.AliasName).IdName(config.IdName).Joins(config.Joins...)

	if config.Debug {
		newRepo.Debug()
	}

	if config.ThreadSafe {
		newRepo.ThreadSafe()
	}

	return newRepo
}
