package update_builder

import (
	"github.com/lowl11/lazy-entity/builders/condition_builder"
	"strings"
)

type Builder struct {
	condition_builder.Builder

	tableName string

	conditions *strings.Builder
	setValues  []string
	setFields  []string
}

func New(tableName string) *Builder {
	builder := &Builder{
		tableName:  tableName,
		conditions: &strings.Builder{},
		setValues:  make([]string, 0),
		setFields:  make([]string, 0),
	}

	builder.Builder = condition_builder.New(builder.getFieldItem)
	return builder
}
