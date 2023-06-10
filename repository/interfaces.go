package repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories"
)

type ICrudRepository[T any, ID repositories.IComparableID] interface {
	Count() (int, error)
	GetAll() ([]T, error)
	GetByID(id ID) (*T, error)
	ExistByID(id ID) (bool, error)
	Add(entity T) (ID, error)
	DeleteAll() error
	DeleteByID(id ID) error
}
