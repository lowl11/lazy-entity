package queryapi

import (
	"github.com/lowl11/lazy-entity/internal/builders/delete_builder"
	"github.com/lowl11/lazy-entity/internal/builders/insert_builder"
	"github.com/lowl11/lazy-entity/internal/builders/select_builder"
	"github.com/lowl11/lazy-entity/internal/builders/update_builder"
)

func Select(fields ...string) *select_builder.Builder {
	return select_builder.New(fields...)
}

func Insert(tableName string) *insert_builder.Builder {
	return insert_builder.New(tableName)
}

func Update() *update_builder.Builder {
	return update_builder.New()
}

func Delete() *delete_builder.Builder {
	return delete_builder.New()
}
