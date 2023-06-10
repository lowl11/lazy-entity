package crud_repository

import (
	"github.com/lowl11/lazy-entity/queryapi"
)

func (repo *CrudRepository[T, ID]) Alias(aliasName string) *CrudRepository[T, ID] {
	repo.aliasName = aliasName
	return repo
}

func (repo *CrudRepository[T, ID]) Count() (int, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Select("count(*)").
		From(repo.tableName).
		Build()

	count := -1
	if err := repo.connection.QueryRowxContext(ctx, query).Scan(&count); err != nil {
		return count, err
	}

	return count, nil
}

func (repo *CrudRepository[T, ID]) GetAll() ([]T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Select(repo.fieldList...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Build()

	rows, err := repo.connection.QueryxContext(ctx, query)
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

func (repo *CrudRepository[T, ID]) GetByID(id ID) (*T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	query := builder.
		Fields(repo.fieldList...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Where(builder.Equal("id", "$1")).
		Build()

	rows, err := repo.connection.QueryxContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer repo.CloseRows(rows)

	if rows.Next() {
		var item T
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (repo *CrudRepository[T, ID]) Add(entity T) (ID, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.fieldListWithoutID()...).
		Returning("id").
		VariableMode().
		Build()

	var id ID
	rows, err := repo.connection.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return id, err
	}

	if !rows.Next() {
		return id, nil
	}

	if err = rows.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (repo *CrudRepository[T, ID]) AddList(entityList []T) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.fieldListWithoutID()...).
		Returning("id").
		VariableMode().
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) DeleteAll() error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Delete(repo.tableName).
		Build()

	if _, err := repo.connection.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) DeleteByID(id ID) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Delete(repo.tableName)
	query := builder.
		Where(builder.Equal("id", id)).
		Build()

	if _, err := repo.connection.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) ExistByID(id ID) (bool, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	query := builder.
		Fields("id").
		From(repo.tableName).
		Where(builder.Equal("id", "$1")).
		Build()

	rows, err := repo.connection.QueryxContext(ctx, query, id)
	if err != nil {
		return false, err
	}
	defer repo.CloseRows(rows)

	return rows.Next(), nil
}
