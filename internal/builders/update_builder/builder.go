package update_builder

import "github.com/lowl11/lazy-entity/internal/builders/condition_builder"

type Builder struct {
	condition_builder.Builder

	tableName string

	conditions string
	setValues  []string
	setFields  []string
}

func New(tableName string) *Builder {
	builder := &Builder{
		tableName: tableName,
		setValues: make([]string, 0),
		setFields: make([]string, 0),
	}

	builder.Builder = condition_builder.New(builder.getFieldItem)
	return builder
}
