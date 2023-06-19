package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
)

type IUniversalRepository[T any, ID repositories.IComparableID] interface {
	Guid() string

	Count() (int, error)
	Exist(customizeFunc func(builder *select_builder.Builder), args ...any) (bool, error)

	GetList(func(builder *select_builder.Builder), ...any) ([]T, error)
	GetItem(func(builder *select_builder.Builder), ...any) (*T, error)

	Add(entity T) (ID, error)
	AddTx(tx *sqlx.Tx, entity T) (ID, error)
	AddWithID(entity T) error
	AddWithIdTx(tx *sqlx.Tx, entity T) error
	AddList(entityList []T) error
	AddListTx(tx *sqlx.Tx, entityList []T) error
	AddListWithID(entityList []T) error
	AddListWithIdTx(tx *sqlx.Tx, entityList []T) error

	Update(
		customizeFunc func(builder *update_builder.Builder),
		entity T,
	) error
	UpdateTx(
		tx *sqlx.Tx,
		customizeFunc func(builder *update_builder.Builder),
		entity T,
	) error

	Delete(customizeFunc func(builder *delete_builder.Builder), args ...any) error
	DeleteTx(tx *sqlx.Tx, customizeFunc func(builder *delete_builder.Builder), args ...any) error

	Transaction(transactionActions func(tx *sqlx.Tx) error) error
}
