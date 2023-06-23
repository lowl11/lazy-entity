package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/crud_repository"
	"github.com/lowl11/lazy-entity/repo_config"
)

// ICrudRepository extends IUniversalRepository and on base of it
// gives basic generally need methods as GetAll()
type ICrudRepository[T any, ID repositories.IComparableID] interface {
	IUniversalRepository[T, ID]

	// ExistByID
	ExistByID(id ID) (bool, error)

	// GetAll
	GetAll() ([]T, error)

	// GetByID
	GetByID(id ID) (*T, error)

	// GetByIdList
	GetByIdList(id []ID) ([]T, error)

	// UpdateByID
	UpdateByID(id ID, entity T) error

	// UpdateByIdTx
	UpdateByIdTx(tx *sqlx.Tx, id ID, entity T) error

	// DeleteAll
	DeleteAll() error

	// DeleteAllTx
	DeleteAllTx(tx *sqlx.Tx) error

	// DeleteByID
	DeleteByID(id ID) error

	// DeleteByIdTx
	DeleteByIdTx(tx *sqlx.Tx, id ID) error
}

func NewCrud[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string, config repo_config.Crud) ICrudRepository[T, ID] {
	// create new instance
	newRepo := crud_repository.New[T, ID](connection, tableName)

	// set given configs
	newRepo.
		Alias(config.AliasName).
		IdName(config.IdName).
		Joins(config.Joins...).
		PageSize(config.PageSize)

	// turn debug mode
	if config.Debug {
		newRepo.Debug()
	}

	// turn thread safe mode
	if config.ThreadSafe {
		newRepo.ThreadSafe()
	}

	return newRepo
}
