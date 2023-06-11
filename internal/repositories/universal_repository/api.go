package universal_repository

import (
	"github.com/lowl11/lazy-entity/internal/builders/select_builder"
	"github.com/lowl11/lazy-entity/queryapi"
)

func (repo *Repository[T, ID]) Alias(aliasName string) *Repository[T, ID] {
	repo.aliasName = aliasName
	return repo
}

func (repo *Repository[T, ID]) IdName(name string) *Repository[T, ID] {
	if name == "" {
		return repo
	}

	repo.idName = name
	return repo
}

func (repo *Repository[T, ID]) GetList(conditionFunc func(builder *select_builder.Builder) string, args ...any) ([]T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	builder.
		Fields(repo.fieldList...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Where(conditionFunc(builder))

	rows, err := repo.connection.QueryxContext(ctx, builder.Build(), args...)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	list := make([]T, 0)
	for rows.Next() {
		var item T
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
