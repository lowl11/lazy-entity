package repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
)

type IUniversalRepository[T any, ID repositories.IComparableID] interface {
	Count() (int, error)
	Exist(customizeFunc func(builder *select_builder.Builder), args ...any) (bool, error)

	GetList(func(builder *select_builder.Builder), ...any) ([]T, error)
	GetItem(func(builder *select_builder.Builder), ...any) (*T, error)

	Add(entity T) (ID, error)
	AddWithID(entity T) error
	AddList(entityList []T) error
	AddListWithID(entityList []T) error

	Update(
		customizeFunc func(builder *update_builder.Builder),
		entity T,
	) error

	Delete(customizeFunc func(builder *delete_builder.Builder), args ...any) error
}
