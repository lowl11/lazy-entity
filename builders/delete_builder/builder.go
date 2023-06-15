package delete_builder

import (
	"github.com/lowl11/lazy-entity/builders/condition_builder"
	"strings"
)

type Builder struct {
	condition_builder.Builder

	tableName  string
	conditions *strings.Builder
}

func New(tableName string) *Builder {
	builder := &Builder{
		tableName:  tableName,
		conditions: &strings.Builder{},
	}

	builder.Builder = condition_builder.New(builder.getFieldItem)
	return builder
}
