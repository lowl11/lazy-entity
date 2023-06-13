package universal_repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/queryapi"
	"github.com/lowl11/lazy-entity/repo_config"
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

func (repo *Repository[T, ID]) Joins(joinList ...repo_config.Join) *Repository[T, ID] {
	repo.joinList = append(repo.joinList, joinList...)
	return repo
}

func (repo *Repository[T, ID]) GetIdName() string {
	return repo.idName
}

func (repo *Repository[T, ID]) ExistByID(id ID) (bool, error) {
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

func (repo *Repository[T, ID]) Count() (int, error) {
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

func (repo *Repository[T, ID]) GetList(customizeFunc func(builder *select_builder.Builder), args ...any) ([]T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	builder.
		Fields(repo.getFieldList(true)...).
		From(repo.tableName).
		Alias(repo.aliasName)

	for _, item := range repo.joinList {
		builder.Join(item.TableName, item.AliasName, item.Condition(builder))
	}

	customizeFunc(builder)

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

func (repo *Repository[T, ID]) GetItem(customizeFunc func(builder *select_builder.Builder), args ...any) (*T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	builder.
		Fields(repo.fieldList...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Limit(1)

	customizeFunc(builder)

	rows, err := repo.connection.QueryxContext(ctx, builder.Build(), args...)
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

func (repo *Repository[T, ID]) Add(entity T) (ID, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFieldList(false)...).
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

func (repo *Repository[T, ID]) AddList(entityList []T) error {
	if len(entityList) == 0 {
		return nil
	}

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFieldList(false)...).
		VariableMode().
		Build()

	if _, err := repo.connection.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) Update(
	conditionFunc func(builder *update_builder.Builder) string,
	entity T,
) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)
	query := builder.
		Where(conditionFunc(builder)).
		Build()

	nonEmptyIndices := type_helper.GetObjectNonEmptyIndices(&entity)
	builder.SetByFields(repo.getNonEmptyFields(nonEmptyIndices)...)

	if _, err := repo.connection.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) Delete(customizeFunc func(builder *delete_builder.Builder), args ...any) error {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Delete(repo.tableName)
	customizeFunc(builder)

	if _, err := repo.connection.ExecContext(ctx, builder.Build(), args...); err != nil {
		return err
	}

	return nil
}
