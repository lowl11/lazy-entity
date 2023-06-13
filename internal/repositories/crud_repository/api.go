package crud_repository

import (
	"github.com/lib/pq"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/order_types"
)

func (repo *Repository[T, ID]) GetAll() ([]T, error) {
	return repo.GetList(func(builder *select_builder.Builder) {})
}

func (repo *Repository[T, ID]) GetByID(id ID) (*T, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal(repo.Repository.GetIdName(), id))
	})
}

func (repo *Repository[T, ID]) GetByIdList(id []ID) ([]T, error) {
	return repo.GetList(func(builder *select_builder.Builder) {
		builder.
			Where(builder.In(repo.GetIdName(), "$1")).
			OrderBy(order_types.Asc, repo.GetIdName())
	}, pq.Array(id))
}

func (repo *Repository[T, ID]) UpdateByID(id ID, entity T) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	}, entity)
}

func (repo *Repository[T, ID]) DeleteAll() error {
	return repo.Delete(func(builder *delete_builder.Builder) {})
}

func (repo *Repository[T, ID]) DeleteByID(id ID) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal(repo.GetIdName(), id))
	})
}
