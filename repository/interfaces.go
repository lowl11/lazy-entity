package repository

import (
	"github.com/lowl11/lazy-entity/internal/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
)

type ICrudRepository[T any, ID repositories.IComparableID] interface {
	Count() (int, error)
	ExistByID(id ID) (bool, error)

	GetAll() ([]T, error)
	GetByID(id ID) (*T, error)

	Add(entity T) (ID, error)
	AddList(entityList []T) error

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
