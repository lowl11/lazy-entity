package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/repositories"
)

// IUniversalRepository gives all possible abstract CRUD methods.
// Some methods can be modified by using "builders"
type IUniversalRepository[T any, ID repositories.IComparableID] interface {
	// Guid returns generated new UUID.
	// Usually, it uses for generating some new entity IDs.
	// This method already given in IRepository interface
	Guid() string

	// Count get quantity of all records of entity.
	// For counting some "specific" amount of records with condition.
	// 	use Select("COUNT(*)").From("table_name").Where(...)
	Count(ctx context.Context, customizeFunc func(builder *select_builder.Builder)) (int, error)

	// Exist returns true if record is found.
	// Without conditions, it will return true if there is at least 1 record
	Exist(ctx context.Context, customizeFunc func(builder *select_builder.Builder), args ...any) (bool, error)

	// GetItem returns 1 entity object with given conditions or NULL (nil)
	GetItem(context.Context, func(builder *select_builder.Builder), ...any) (*T, error)

	// GetList returns list of entity objects with given conditions or empty slice.
	GetList(context.Context, func(builder *select_builder.Builder), ...any) ([]T, error)

	// GetPage returns list of entity objects with given conditions and given page or empty slice.
	// By default, there is pageSize=10. But it is mutable with repo_config.Crud{}
	GetPage(context.Context, int, func(builder *select_builder.Builder), ...any) ([]T, error)

	// Add new entity with ignoring ID field (AutoIncrement).
	// Inserting field list depends on which fields were fill and which not
	Add(ctx context.Context, entity T) (ID, error)

	// AddTx the same as Add but using transaction
	AddTx(ctx context.Context, tx *sqlx.Tx, entity T) (ID, error)

	// AddWithID new entity, with ID (no AutoIncrement)
	AddWithID(ctx context.Context, entity T) error

	// AddWithIdTx same as AddWithID but using transaction
	AddWithIdTx(ctx context.Context, tx *sqlx.Tx, entity T) error

	// AddList creates list of new records with generated IDs
	AddList(ctx context.Context, entityList []T) error

	// AddListTx the same as AddList but using transaction
	AddListTx(ctx context.Context, tx *sqlx.Tx, entityList []T) error

	// AddListWithID creates list of new records (NoAutoincrement)
	AddListWithID(ctx context.Context, entityList []T) error

	// AddListWithIdTx the same as AddListWithID but using transaction
	AddListWithIdTx(ctx context.Context, tx *sqlx.Tx, entityList []T) error

	// Update updates record with given entity object.
	// Set fields depends on which object fields are fill and which not
	Update(
		ctx context.Context,
		customizeFunc func(builder *update_builder.Builder),
		entity T,
	) error

	// UpdateTx the same as Update but using transaction
	UpdateTx(
		ctx context.Context,
		tx *sqlx.Tx,
		customizeFunc func(builder *update_builder.Builder),
		entity T,
	) error

	// Delete removes record with given conditions
	Delete(ctx context.Context, customizeFunc func(builder *delete_builder.Builder), args ...any) error

	// DeleteTx the same as Delete but using transaction
	DeleteTx(ctx context.Context, tx *sqlx.Tx, customizeFunc func(builder *delete_builder.Builder), args ...any) error

	// Transaction takes func as an argument which will take transaction object.
	// Also, there is one more way to use transaction by calling "transaction_service"
	Transaction(transactionActions func(tx *sqlx.Tx) error) error
}
