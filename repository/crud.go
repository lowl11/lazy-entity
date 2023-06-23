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

	// ExistByID returns true if record exist
	ExistByID(id ID) (bool, error)

	// GetAll returns all records
	GetAll() ([]T, error)

	// GetByID returns record by ID. 
	// If the record with such ID does not exist, returns NULL (nil)
	GetByID(id ID) (*T, error)

	// GetByIdList returns list of records by array of IDs.
	// Some of IDs could be not existing, there wouldn't be an error
	GetByIdList(id []ID) ([]T, error)

	// UpdateByID updates an record by ID
	UpdateByID(id ID, entity T) error

	// UpdateByIdTx the same as UpdateByID but using transaction
	UpdateByIdTx(tx *sqlx.Tx, id ID, entity T) error

	// DeleteAll remove all records
	DeleteAll() error

	// DeleteAllTx the same as DeleteAll but using transaction
	DeleteAllTx(tx *sqlx.Tx) error

	// DeleteByID removes record by given ID
	DeleteByID(id ID) error

	// DeleteByIdTx the same as DeleteByID but using transaction
	DeleteByIdTx(tx *sqlx.Tx, id ID) error
}

func NewCrud[T any, ID repositories.IComparableID](
	connection *sqlx.DB,
	tableName string,
	config repo_config.Crud,
) ICrudRepository[T, ID] {
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
