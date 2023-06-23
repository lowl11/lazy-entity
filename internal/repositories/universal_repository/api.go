package universal_repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"github.com/lowl11/lazy-entity/queryapi"
	"github.com/lowl11/lazy-entity/repo_config"
	"sync"
)

func (repo *Repository[T, ID]) Guid() string {
	return repo.Repository.Guid()
}

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

func (repo *Repository[T, ID]) PageSize(size int) *Repository[T, ID] {
	repo.pageSize = size
	return repo
}

func (repo *Repository[T, ID]) Joins(joinList ...repo_config.Join) *Repository[T, ID] {
	repo.joinList = append(repo.joinList, joinList...)
	return repo
}

func (repo *Repository[T, ID]) ThreadSafe() *Repository[T, ID] {
	repo.threadSafe = true
	repo.mutex = &sync.Mutex{}
	return repo
}

func (repo *Repository[T, ID]) GetIdName() string {
	return repo.idName
}

func (repo *Repository[T, ID]) Debug() *Repository[T, ID] {
	repo.debug = true
	return repo
}

func (repo *Repository[T, ID]) Exist(customizeFunc func(builder *select_builder.Builder), args ...any) (bool, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	builder.
		Fields(repo.idName).
		From(repo.tableName)

	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	rows, err := repo.connection.QueryxContext(ctx, query, args...)
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

	if repo.debug {
		fmt.Println(query)
	}

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
		Fields(repo.getFieldsWithJoin(true)...).
		From(repo.tableName).
		Alias(repo.aliasName)

	for _, item := range repo.joinList {
		builder.Join(item.TableName, item.AliasName, item.Condition(builder))
	}

	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	rows, err := repo.connection.QueryxContext(ctx, query, args...)
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

func (repo *Repository[T, ID]) GetPage(pageNum int, customizeFunc func(builder *select_builder.Builder), args ...any) ([]T, error) {
	return repo.GetList(func(builder *select_builder.Builder) {
		customizeFunc(builder)

		pageSize := repo.pageSize

		// if page size is not given
		if pageSize == 0 {
			pageSize = 10
		}

		builder.
			Offset(pageNum * pageSize).
			Limit(pageSize)
	})
}

func (repo *Repository[T, ID]) GetItem(customizeFunc func(builder *select_builder.Builder), args ...any) (*T, error) {
	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Select()
	builder.
		Fields(repo.getFieldsWithJoin(true)...).
		From(repo.tableName).
		Alias(repo.aliasName).
		Limit(1)

	for _, item := range repo.joinList {
		builder.Join(item.TableName, item.AliasName, item.Condition(builder))
	}

	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	rows, err := repo.connection.QueryxContext(ctx, query, args...)
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
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(false)...).
		Returning(repo.idName).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

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

func (repo *Repository[T, ID]) AddTx(tx *sqlx.Tx, entity T) (ID, error) {
	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(false)...).
		Returning(repo.idName).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	var id ID
	rows, err := tx.NamedQuery(query, entity)
	if err != nil {
		return id, err
	}

	if !rows.Next() {
		return id, nil
	}

	if err = rows.Scan(&id); err != nil {
		return id, err
	}
	repo.CloseRows(rows)

	return id, nil
}

func (repo *Repository[T, ID]) AddWithID(entity T) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(true)...).
		Returning(repo.idName).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := repo.connection.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) AddWithIdTx(tx *sqlx.Tx, entity T) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(true)...).
		Returning(repo.idName).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) AddList(entityList []T) error {
	if len(entityList) == 0 {
		return nil
	}

	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(false)...).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := repo.connection.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) AddListTx(tx *sqlx.Tx, entityList []T) error {
	if len(entityList) == 0 {
		return nil
	}

	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(false)...).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := tx.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) AddListWithID(entityList []T) error {
	if len(entityList) == 0 {
		return nil
	}

	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(true)...).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := repo.connection.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) AddListWithIdTx(tx *sqlx.Tx, entityList []T) error {
	if len(entityList) == 0 {
		return nil
	}

	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	query := queryapi.
		Insert(repo.tableName).
		Fields(repo.getFields(true)...).
		VariableMode().
		Build()

	if repo.debug {
		fmt.Println(query)
	}

	if _, err := tx.NamedExecContext(ctx, query, entityList); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) Update(
	customizeFunc func(builder *update_builder.Builder),
	entity T,
) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)

	nonEmptyIndices := type_helper.GetObjectNonEmptyIndices(&entity)
	nonEmptyFields := repo.getNonEmptyFields(nonEmptyIndices)
	if len(nonEmptyFields) == 0 {
		return nil
	}

	builder.SetByFields(nonEmptyFields...)
	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	if _, err := repo.connection.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) UpdateTx(
	tx *sqlx.Tx,
	customizeFunc func(builder *update_builder.Builder),
	entity T,
) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Update(repo.tableName)

	nonEmptyIndices := type_helper.GetObjectNonEmptyIndices(&entity)
	nonEmptyFields := repo.getNonEmptyFields(nonEmptyIndices)
	if len(nonEmptyFields) == 0 {
		return nil
	}

	builder.SetByFields(nonEmptyFields...)
	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	if _, err := tx.NamedExecContext(ctx, query, entity); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) Delete(customizeFunc func(builder *delete_builder.Builder), args ...any) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Delete(repo.tableName)
	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	if _, err := repo.connection.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) DeleteTx(tx *sqlx.Tx, customizeFunc func(builder *delete_builder.Builder), args ...any) error {
	repo.lock()
	defer repo.unlock()

	ctx, cancel := repo.Ctx()
	defer cancel()

	builder := queryapi.Delete(repo.tableName)
	customizeFunc(builder)

	query := builder.Build()
	if repo.debug {
		fmt.Println(query)
	}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (repo *Repository[T, ID]) Transaction(transactionActions func(tx *sqlx.Tx) error) error {
	repo.lock()
	defer repo.unlock()

	return repo.Repository.Transaction(repo.connection, func(tx *sqlx.Tx) error {
		return transactionActions(tx)
	})
}
