package queryapi

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/insert_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
)

func Select(tableName string) *select_builder.Builder {
	return select_builder.
		New(tableName)
}

func Insert(tableName string) *insert_builder.Builder {
	return insert_builder.
		New(tableName)
}

func Update(tableName string) *update_builder.Builder {
	return update_builder.
		New(tableName)
}

func Delete(tableName string) *delete_builder.Builder {
	return delete_builder.
		New(tableName)
}
