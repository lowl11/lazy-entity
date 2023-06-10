package crud_repository

import (
	"github.com/lib/pq"
	"github.com/lowl11/lazy-entity/internal/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/queryapi"
)

func (repo *CrudRepository[T, ID]) Alias(aliasName string) *CrudRepository[T, ID] {
	repo.aliasName = aliasName
	return repo
}

func (repo *CrudRepository[T, ID]) IdName(name string) *CrudRepository[T, ID] {
	if name == "" {
		return repo
	}

	repo.idName = name
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

func (repo *CrudRepository[T, ID]) ExistByID(id ID) (bool, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	query := builder.
		Fields(repo.idName).
		From(repo.tableName).
		Where(builder.Equal(repo.idName, "$1")).
		Build()

	rows, err := repo.connection.QueryxContext(ctx, query, id)
	if err != nil {
		return false, err
	}
	defer repo.CloseRows(rows)

	return rows.Next(), nil
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
		Where(builder.Equal(repo.idName, "$1")).
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

func (repo *CrudRepository[T, ID]) GetByIdList(id []ID) ([]T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	query := builder.
		Fields(repo.fieldList...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Where(builder.In(repo.idName, "$1")).
		Build()

	rows, err := repo.connection.QueryxContext(ctx, query, pq.Array(id))
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

func (repo *CrudRepository[T, ID]) Add(entity any) (ID, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	fieldList := type_helper.GetStructFieldsByObject(entity)

	query := queryapi.
		Insert(repo.tableName).
		Fields(fieldList...).
		Returning(repo.idName).
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

func (repo *CrudRepository[T, ID]) AddList(entityList []any) error {
	if len(entityList) == 0 {
		return nil
	}

	ctx, cancel := repo.Ctx()
	defer cancel()

	fieldList := type_helper.GetStructFieldsByObject(entityList[0])

	query := queryapi.
		Insert(repo.tableName).
		Fields(fieldList...).
		Returning(repo.idName).
		VariableMode().
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) SaveByID(id ID, entity T) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)
	query := builder.
		SetByFields(repo.fieldList...).
		Where(builder.Equal(repo.idName, id)).
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) SaveByCondition(
	conditionFunc func(builder *update_builder.Builder) string,
	entity T,
) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)
	query := builder.
		SetByFields(repo.fieldList...).
		Where(conditionFunc(builder)).
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) UpdateByID(id ID, updateEntity any) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)
	query := builder.
		SetByFields(type_helper.GetStructFieldsByObject(updateEntity)...).
		Where(builder.Equal(repo.idName, id)).
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, updateEntity); err != nil {
		return err
	}

	return nil
}

func (repo *CrudRepository[T, ID]) UpdateByCondition(
	conditionFunc func(builder *update_builder.Builder) string,
	updateEntity any,
) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)
	query := builder.
		SetByFields(type_helper.GetStructFieldsByObject(updateEntity)...).
		Where(conditionFunc(builder)).
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, updateEntity); err != nil {
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
		Where(builder.Equal(repo.idName, id)).
		Build()

	if _, err := repo.connection.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
