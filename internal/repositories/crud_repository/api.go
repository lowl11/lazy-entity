package crud_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/enums/order_types"
)

func (repo *Repository[T, ID]) CountAll(ctx context.Context) (int, error) {
	return repo.Count(ctx, func(builder *select_builder.Builder) {})
}

func (repo *Repository[T, ID]) ExistByID(ctx context.Context, id ID) (bool, error) {
	return repo.Exist(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), "$1"))
	}, id)
}

func (repo *Repository[T, ID]) GetAll(ctx context.Context) ([]T, error) {
	return repo.GetList(ctx, func(builder *select_builder.Builder) {})
}

func (repo *Repository[T, ID]) GetByID(ctx context.Context, id ID) (*T, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Equal(repo.Repository.GetIdName(), id))
	})
}

func (repo *Repository[T, ID]) GetByIdList(ctx context.Context, id []ID) ([]T, error) {
	return repo.GetList(ctx, func(builder *select_builder.Builder) {
		builder.
			Where(builder.In(repo.GetIdName(), "$1")).
			OrderBy(order_types.Asc, repo.GetIdName())
	}, pq.Array(id))
}

func (repo *Repository[T, ID]) UpdateByID(ctx context.Context, id ID, entity T) error {
	return repo.Update(ctx, func(builder *update_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	}, entity)
}

func (repo *Repository[T, ID]) UpdateByIdTx(ctx context.Context, tx *sqlx.Tx, id ID, entity T) error {
	return repo.UpdateTx(ctx, tx, func(builder *update_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	}, entity)
}

func (repo *Repository[T, ID]) DeleteAll(ctx context.Context) error {
	return repo.Delete(ctx, func(builder *delete_builder.Builder) {})
}

func (repo *Repository[T, ID]) DeleteAllTx(ctx context.Context, tx *sqlx.Tx) error {
	return repo.DeleteTx(ctx, tx, func(builder *delete_builder.Builder) {})
}

func (repo *Repository[T, ID]) DeleteByID(ctx context.Context, id ID) error {
	return repo.Delete(ctx, func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	})
}

func (repo *Repository[T, ID]) DeleteByIdTx(ctx context.Context, tx *sqlx.Tx, id ID) error {
	return repo.DeleteTx(ctx, tx, func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	})
}
